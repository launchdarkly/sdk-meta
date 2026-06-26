FROM eclipse-temurin:17-jdk-noble

# Android SDK command-line tools + the build-tools / platforms versions
# the snippet needs. We *don't* install the emulator: Robolectric runs
# the activity lifecycle in the JVM, which is enough to exercise the
# snippet's init+evaluate path against the real LD streaming API.
# apt-get update against the noble repos is occasionally flaky on CI
# build hosts (intermittent timeouts to security.ubuntu.com). Retry
# both the index update and the install with explicit acquire-retry
# tuning before giving up so a single fetch failure doesn't fail the
# whole image build.
RUN apt-get update -o Acquire::Retries=5 \
        -o Acquire::http::Timeout=30 \
        -o Acquire::https::Timeout=30 \
    && apt-get install -y --no-install-recommends \
        -o Acquire::Retries=5 \
        curl unzip git ca-certificates python3-minimal \
    && rm -rf /var/lib/apt/lists/*

ENV ANDROID_HOME=/opt/android-sdk
ENV ANDROID_SDK_ROOT=${ANDROID_HOME}
ENV PATH="${ANDROID_HOME}/cmdline-tools/latest/bin:${ANDROID_HOME}/platform-tools:${PATH}"

# Pin a specific cmdline-tools build to avoid surprise upgrades.
ARG CMDLINE_TOOLS_VERSION=11076708
RUN mkdir -p ${ANDROID_HOME}/cmdline-tools \
    && curl -fsSL --retry 3 --retry-delay 5 --max-time 600 "https://dl.google.com/android/repository/commandlinetools-linux-${CMDLINE_TOOLS_VERSION}_latest.zip" -o /tmp/tools.zip \
    && unzip -q /tmp/tools.zip -d ${ANDROID_HOME}/cmdline-tools \
    && mv ${ANDROID_HOME}/cmdline-tools/cmdline-tools ${ANDROID_HOME}/cmdline-tools/latest \
    && rm /tmp/tools.zip

RUN yes | sdkmanager --licenses >/dev/null \
    && sdkmanager --install "platforms;android-35" "platform-tools" "build-tools;35.0.0"

# Pre-bake a hello-android gradle scaffold. We clone the upstream
# hello-android repo (LD's official getting-started reference) at a
# pinned ref, then patch app/build.gradle to use the latest SDK +
# Robolectric (so the validator runs in a pure JVM, no emulator).
ARG HELLO_ANDROID_REF=main
# 5.13.0 introduced the EAP data system surface (Components.dataSystem(),
# DataSystemBuilder, AutomaticModeSwitchingConfig) that the
# data-saving-mode doc fragments compile against.
ARG LD_ANDROID_SDK_VERSION=5.13.1
ARG ROBOLECTRIC_VERSION=4.12.2
# Optional observability plugin AAR. Doc fragments under
# initialize-the-client-android-sdk-v5-x-* show the v5.9+ pattern
# `Components.plugins().setPlugins(listOf(Observability(this@…)))`,
# which references the launchdarkly-observability-android package.
ARG LD_ANDROID_OBSERVABILITY_VERSION=0.54.0

RUN git clone --depth 1 --branch "${HELLO_ANDROID_REF}" \
        https://github.com/launchdarkly/hello-android.git /opt/hello-android

WORKDIR /opt/hello-android

# Patch the LD SDK version, add the observability AAR and Timber
# (referenced by log-statement doc fragments) to `implementation`,
# and add Robolectric to `testImplementation`.
# Also bump compileSdk/targetSdk to 35 because the v0.49.0
# observability AAR transitively depends on androidx.core 1.16+,
# which requires API 35.
RUN sed -i "s|com.launchdarkly:launchdarkly-android-client-sdk:[0-9.]*|com.launchdarkly:launchdarkly-android-client-sdk:${LD_ANDROID_SDK_VERSION}|" app/build.gradle \
    && sed -i "s|compileSdk 34|compileSdk 35|; s|targetSdkVersion(34)|targetSdkVersion(35)|" app/build.gradle \
    && sed -i "/^dependencies {$/a\\    implementation 'com.launchdarkly:launchdarkly-observability-android:${LD_ANDROID_OBSERVABILITY_VERSION}'\\n    implementation 'com.jakewharton.timber:timber:5.0.1'\\n    testImplementation 'org.robolectric:robolectric:${ROBOLECTRIC_VERSION}'\\n    testImplementation 'junit:junit:4.13.2'\\n    testImplementation 'androidx.test:core:1.5.0'\\n    testImplementation 'androidx.test.ext:junit:1.1.5'" app/build.gradle

# Robolectric needs testOptions { unitTests.includeAndroidResources } to
# resolve R.string / R.id at test time. Forward stdout/stderr so the
# harness can grep the canonical EXAM-HELLO line; without this the
# textView dump only ends up in build/reports/tests.
RUN cat >> app/build.gradle <<'EOF'

android {
    testOptions {
        unitTests {
            includeAndroidResources = true
            all {
                testLogging {
                    events 'passed', 'failed', 'standardOut', 'standardError'
                    showStandardStreams = true
                }
            }
        }
    }
}
EOF

# Drop a Robolectric test that drives the Activity lifecycle and reads
# the textView's content. This file is harness-internal; the snippet
# itself stays unchanged.
RUN mkdir -p app/src/test/java/com/launchdarkly/hello_android
COPY languages/android-client/test/HelloAppTest.kt \
        app/src/test/java/com/launchdarkly/hello_android/HelloAppTest.kt

# Pre-warm: pull deps, compile, run a placeholder test cycle.
RUN ./gradlew --no-daemon dependencies --console=plain >/dev/null 2>&1 || true

WORKDIR /work

COPY shared /harness-shared
COPY languages/android-client/harness /harness

ENTRYPOINT ["/harness/run.sh"]
