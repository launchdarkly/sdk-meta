#!/bin/sh
# Runs an iOS install-fragment snippet on a clean macOS runner.
#
# Native-mode harness: the Go validator stages the snippet under
# $SNIPPET_DIR and execs us directly with $SNIPPET_ENTRYPOINT pointing
# at the (single-file) install fragment. Each fragment kind is wrapped
# in a minimal project before the package manager runs.
#
# Inputs (env):
#   SNIPPET_DIR        — staged snippet root.
#   SNIPPET_ENTRYPOINT — path (relative to $SNIPPET_DIR) of the install
#                        fragment body.
#   INSTALL_KIND       — swift-package | podfile | cartfile.
set -eu

if [ -z "${SNIPPET_DIR:-}" ] || [ -z "${SNIPPET_ENTRYPOINT:-}" ]; then
    echo "ios-install: SNIPPET_DIR and SNIPPET_ENTRYPOINT must be set" >&2
    exit 1
fi

KIND="${INSTALL_KIND:-}"
if [ -z "$KIND" ]; then
    echo "ios-install: INSTALL_KIND not set" >&2
    exit 1
fi

BODY_FILE="$SNIPPET_DIR/$SNIPPET_ENTRYPOINT"
if [ ! -f "$BODY_FILE" ]; then
    echo "ios-install: install fragment not found at $BODY_FILE" >&2
    exit 1
fi

WORK=$(mktemp -d 2>/dev/null || mktemp -d -t ios-install.XXXXXX)
trap 'rm -rf "$WORK"' EXIT
cd "$WORK"

case "$KIND" in
    swift-package)
        # Wrap the snippet's `dependencies:` / `targets:` arrays in a
        # minimal Package.swift skeleton, substitute the snippet's
        # `YOUR_TARGET` placeholder for a fixture target, then resolve.
        # Assert Package.resolved references the LaunchDarkly SDK
        # repository so a snippet that lost the SDK declaration would
        # fail.
        mkdir -p Sources/HelloLDFixture
        cat > Sources/HelloLDFixture/main.swift <<'EOF'
print("hello-ld-fixture")
EOF
        BODY=$(sed 's/YOUR_TARGET/HelloLDFixture/g' "$BODY_FILE")
        cat > Package.swift <<EOF
// swift-tools-version:5.7
import PackageDescription

let package = Package(
    name: "HelloLDFixture",
${BODY}
)
EOF
        if ! swift package resolve >resolve.log 2>&1; then
            echo "ios-install: swift package resolve failed" >&2
            cat resolve.log >&2
            exit 1
        fi
        if [ ! -f Package.resolved ]; then
            echo "ios-install: Package.resolved was not created" >&2
            cat resolve.log >&2
            exit 1
        fi
        if ! grep -q "ios-client-sdk" Package.resolved; then
            echo "ios-install: Package.resolved does not reference launchdarkly/ios-client-sdk" >&2
            cat Package.resolved >&2
            exit 1
        fi
        echo "validator: ok — Package.resolved references launchdarkly/ios-client-sdk"
        ;;

    podfile)
        # The snippet is a `target 'YourTargetName' do ... pod
        # 'LaunchDarkly' ... end` fragment. We need a dummy Xcode
        # project that owns a target with that exact name for `pod
        # install` to succeed. xcodeproj's gem can synthesize one
        # without invoking xcodebuild.
        if ! command -v pod >/dev/null 2>&1; then
            echo "ios-install: cocoapods (pod) not found on PATH" >&2
            echo "ios-install: install with `gem install cocoapods` on the runner" >&2
            exit 1
        fi
        if ! command -v ruby >/dev/null 2>&1; then
            echo "ios-install: ruby not found on PATH" >&2
            exit 1
        fi

        # The snippet's hardcoded target name. Stay verbatim — that's
        # what we're validating.
        TARGET_NAME='YourTargetName'

        cat > make_project.rb <<'EOF'
require 'xcodeproj'
project = Xcodeproj::Project.new('YourTargetName.xcodeproj')
target = project.new_target(:application, 'YourTargetName', :ios, '13.0')
project.save
EOF
        if ! ruby -rxcodeproj make_project.rb >project.log 2>&1; then
            echo "ios-install: xcodeproj synthesize failed" >&2
            cat project.log >&2
            echo "ios-install: hint — `gem install xcodeproj` on the runner" >&2
            exit 1
        fi

        cp "$BODY_FILE" Podfile

        if ! pod install --no-repo-update >pod.log 2>&1; then
            echo "ios-install: pod install failed" >&2
            cat pod.log >&2
            exit 1
        fi

        if [ ! -d "Pods/LaunchDarkly" ]; then
            echo "ios-install: Pods/LaunchDarkly directory not found after pod install" >&2
            ls -la Pods/ >&2 || true
            exit 1
        fi
        echo "validator: ok — Pods/LaunchDarkly present after pod install"
        ;;

    cartfile)
        if ! command -v carthage >/dev/null 2>&1; then
            echo "ios-install: carthage not found on PATH" >&2
            exit 1
        fi
        cp "$BODY_FILE" Cartfile
        # --no-build: just resolve + check out, don't try to xcodebuild.
        # --use-xcframeworks is the modern path; harmless even when
        # build is skipped.
        if ! carthage update --use-xcframeworks --no-build >carthage.log 2>&1; then
            echo "ios-install: carthage update failed" >&2
            cat carthage.log >&2
            exit 1
        fi
        if [ ! -f Cartfile.resolved ]; then
            echo "ios-install: Cartfile.resolved was not created" >&2
            cat carthage.log >&2
            exit 1
        fi
        if ! grep -q "ios-client-sdk" Cartfile.resolved; then
            echo "ios-install: Cartfile.resolved does not reference launchdarkly/ios-client-sdk" >&2
            cat Cartfile.resolved >&2
            exit 1
        fi
        echo "validator: ok — Cartfile.resolved references launchdarkly/ios-client-sdk"
        ;;

    *)
        echo "ios-install: unknown INSTALL_KIND: $KIND" >&2
        exit 1
        ;;
esac
