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

# The gonfalon snippet says "Remove the prepopulated lines except the
# first line" — the "first line" is `package com.launchdarkly.tutorial;`
# which `mvn archetype:generate` writes for the user. The snippet itself
# doesn't carry it (gonfalon's UI is showing only the body to type), so
# we add it back here so javac can resolve the mainClass declared in
# the synthesized pom.xml.
appfile="$SNIPPET_ENTRYPOINT"
if [ -f "$appfile" ] && ! head -1 "$appfile" | grep -q '^package '; then
    tmp=$(mktemp)
    printf 'package com.launchdarkly.tutorial;\n\n' > "$tmp"
    cat "$appfile" >> "$tmp"
    mv "$tmp" "$appfile"
fi

cat > pom.xml <<'EOF'
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
           gonfalon snippet shows ${version} fetched from Maven Central;
           we don't reach out from inside the harness. -->
      <version>7.13.4</version>
    </dependency>
  </dependencies>
  <build>
    <plugins>
      <plugin>
        <artifactId>maven-assembly-plugin</artifactId>
        <configuration>
          <archive>
            <manifest>
              <mainClass>com.launchdarkly.tutorial.App</mainClass>
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
