FROM eclipse-temurin:17-jdk-noble

# Android SDK command-line tools + the build-tools / platforms versions
# the snippet needs. We *don't* install the emulator: Robolectric runs
# the activity lifecycle in the JVM, which is enough to exercise the
# snippet's init+evaluate path against the real LD streaming API.
RUN apt-get update && apt-get install -y --no-install-recommends \
        curl unzip git ca-certificates \
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
    && sdkmanager --install "platforms;android-34" "platform-tools" "build-tools;34.0.0"

# Pre-bake a hello-android gradle scaffold. We clone the upstream
# hello-android repo (LD's official getting-started reference) at a
# pinned ref, then patch app/build.gradle to use the latest SDK +
# Robolectric (so the validator runs in a pure JVM, no emulator).
ARG HELLO_ANDROID_REF=main
ARG LD_ANDROID_SDK_VERSION=5.11.1
ARG ROBOLECTRIC_VERSION=4.12.2

RUN git clone --depth 1 --branch "${HELLO_ANDROID_REF}" \
        https://github.com/launchdarkly/hello-android.git /opt/hello-android

WORKDIR /opt/hello-android

# Patch the LD SDK version + add Robolectric to testImplementation.
RUN sed -i "s|com.launchdarkly:launchdarkly-android-client-sdk:[0-9.]*|com.launchdarkly:launchdarkly-android-client-sdk:${LD_ANDROID_SDK_VERSION}|" app/build.gradle \
    && sed -i "/^dependencies {$/a\\    testImplementation 'org.robolectric:robolectric:${ROBOLECTRIC_VERSION}'\\n    testImplementation 'junit:junit:4.13.2'\\n    testImplementation 'androidx.test:core:1.5.0'\\n    testImplementation 'androidx.test.ext:junit:1.1.5'" app/build.gradle

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
