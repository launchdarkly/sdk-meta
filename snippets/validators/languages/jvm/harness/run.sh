#!/bin/sh
# Runs the staged Java snippet against a real LaunchDarkly environment.
# Synthesizes a complete pom.xml around the snippet's App.java rather
# than reproducing gonfalon's `mvn archetype:generate + manual fragment
# pasting` flow — a developer following the gonfalon instructions ends
# up with the same project shape, just authored manually.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_SDK_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cp -r /snippet/. "$WORK/"
cd "$WORK"

# Discover the entrypoint's package + class so the synthesized pom.xml
# can wire mainClass to whatever the staged file declares. Gonfalon's
# original "Remove the prepopulated lines except the first line" hint
# implies `package com.launchdarkly.tutorial;` + class `App`; the
# scaffold-based callers (java-syntax-only) declare their own package
# and class. The harness reads both from the file rather than hard-
# coding either.
appfile="$SNIPPET_ENTRYPOINT"
if [ ! -f "$appfile" ]; then
    echo "validator: snippet entrypoint not found: $appfile" >&2
    exit 1
fi

# If the staged file contains the IMPORT_LIFT_MARKER comment, lift any
# `import …;` lines that appear after it (i.e. inside the wrappee body)
# up to the marker line. Java doesn't permit imports inside method
# bodies, so doc fragments that show install-time imports would
# otherwise fail compilation. The lift is line-based and idempotent.
if grep -q 'IMPORT_LIFT_MARKER' "$appfile"; then
    python3 - "$appfile" <<'PYEOF'
import sys
fp = sys.argv[1]
with open(fp) as f:
    lines = f.read().splitlines()
marker_idx = next((i for i, l in enumerate(lines) if 'IMPORT_LIFT_MARKER' in l), None)
if marker_idx is None:
    sys.exit(0)
# Walk from marker downwards, collect any line whose lstrip() begins
# with `import ` (and ends with `;` to avoid pulling Python-style
# `import_x` identifiers). Replace each with a blank line in the body
# and inject a deduplicated block above the marker.
imports = []
for i in range(marker_idx + 1, len(lines)):
    stripped = lines[i].lstrip()
    if stripped.startswith('import ') and stripped.rstrip().endswith(';'):
        imp = stripped.rstrip()
        if imp not in imports:
            imports.append(imp)
        lines[i] = ''
if imports:
    lines = lines[:marker_idx] + imports + lines[marker_idx:]
with open(fp, 'w') as f:
    f.write('\n'.join(lines) + '\n')
PYEOF
fi

if ! head -1 "$appfile" | grep -q '^package '; then
    # Bodies that don't declare a package are presumed to want the
    # tutorial layout — same as before.
    tmp=$(mktemp)
    printf 'package com.launchdarkly.tutorial;\n\n' > "$tmp"
    cat "$appfile" >> "$tmp"
    mv "$tmp" "$appfile"
fi

PKG=$(grep -m1 '^package ' "$appfile" | sed -e 's/^package //' -e 's/;.*//')
CLS=$(basename "$appfile" .java)
MAIN_CLASS="$PKG.$CLS"

cat > pom.xml <<EOF
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0">
  <modelVersion>4.0.0</modelVersion>
  <groupId>com.launchdarkly.tutorial</groupId>
  <artifactId>hello-java</artifactId>
  <version>1.0-SNAPSHOT</version>
  <properties>
    <maven.compiler.source>17</maven.compiler.source>
    <maven.compiler.target>17</maven.compiler.target>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
  </properties>
  <dependencies>
    <dependency>
      <groupId>com.launchdarkly</groupId>
      <artifactId>launchdarkly-java-server-sdk</artifactId>
      <!-- Pin a recent release so the validator is deterministic. The
           gonfalon snippet shows \${version} fetched from Maven Central;
           we don't reach out from inside the harness. -->
      <version>7.13.4</version>
    </dependency>
    <dependency>
      <groupId>com.launchdarkly</groupId>
      <artifactId>launchdarkly-java-server-sdk-redis-store</artifactId>
      <!-- Store-integration fragments (big segments, persistent feature
           stores) reference the Redis integration's classes; pinned for
           the same determinism reason as the SDK itself. -->
      <version>3.1.1</version>
    </dependency>
    <dependency>
      <groupId>com.launchdarkly</groupId>
      <artifactId>launchdarkly-java-server-sdk-dynamodb-store</artifactId>
      <!-- Persistent-feature-store fragments reference the DynamoDB
           integration's classes; pinned for determinism. -->
      <version>5.0.0</version>
    </dependency>
    <dependency>
      <groupId>com.launchdarkly</groupId>
      <artifactId>launchdarkly-java-server-sdk-consul-store</artifactId>
      <!-- Persistent-feature-store fragments reference the Consul
           integration's classes; pinned for determinism. -->
      <version>5.0.0</version>
    </dependency>
  </dependencies>
  <build>
    <plugins>
      <plugin>
        <artifactId>maven-assembly-plugin</artifactId>
        <configuration>
          <archive>
            <manifest>
              <mainClass>${MAIN_CLASS}</mainClass>
            </manifest>
          </archive>
          <descriptorRefs>
            <descriptorRef>jar-with-dependencies</descriptorRef>
          </descriptorRefs>
        </configuration>
      </plugin>
    </plugins>
  </build>
</project>
EOF

LOG=$(mktemp)
BUILDLOG=$(mktemp)

# Compile + assemble. We keep mvn output in BUILDLOG and only print it
# when the build fails, so a clean run is quiet.
if ! mvn -B -q clean compile assembly:single -DskipTests >"$BUILDLOG" 2>&1; then
    echo "validator: maven build failed" >&2
    echo "--- mvn output ---" >&2
    cat "$BUILDLOG" >&2
    exit 1
fi
rm -f "$BUILDLOG"

CI=1 timeout --signal=TERM 60s java -jar "target/hello-java-1.0-SNAPSHOT-jar-with-dependencies.jar" >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 50 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
