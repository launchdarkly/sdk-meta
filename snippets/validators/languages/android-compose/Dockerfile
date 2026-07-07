# Build context is `validators/`. See validators/languages/python/Dockerfile.
#
# Dedicated type-check validator for the Android observability plugin's
# custom-masking doc fragments (Jetpack Compose `Modifier.ldMask()` /
# `ldUnmask()` and the View `EditText.ldMask()`). These need Jetpack
# Compose, and the plugin's Compose masking API is built against
# androidx.compose 1.7.x, which carries Kotlin 1.9/2.0 metadata — so it
# can't be compiled by the Kotlin-1.8 `android-client` validator. Rather
# than bump that shared validator (and risk every existing android
# snippet), this is a separate Kotlin-2.0 + Compose project. Type-check
# only via `compileDebugKotlin` — no emulator, no run.
FROM eclipse-temurin:17-jdk-noble

RUN apt-get update -o Acquire::Retries=5 \
        -o Acquire::http::Timeout=30 -o Acquire::https::Timeout=30 \
    && apt-get install -y --no-install-recommends -o Acquire::Retries=5 \
        curl unzip git ca-certificates \
    && rm -rf /var/lib/apt/lists/*

ENV ANDROID_HOME=/opt/android-sdk
ENV ANDROID_SDK_ROOT=${ANDROID_HOME}
ENV PATH="${ANDROID_HOME}/cmdline-tools/latest/bin:${ANDROID_HOME}/platform-tools:/opt/gradle/gradle-8.9/bin:${PATH}"

ARG CMDLINE_TOOLS_VERSION=11076708
RUN mkdir -p ${ANDROID_HOME}/cmdline-tools \
    && curl -fsSL --retry 3 --retry-delay 5 --max-time 600 "https://dl.google.com/android/repository/commandlinetools-linux-${CMDLINE_TOOLS_VERSION}_latest.zip" -o /tmp/tools.zip \
    && unzip -q /tmp/tools.zip -d ${ANDROID_HOME}/cmdline-tools \
    && mv ${ANDROID_HOME}/cmdline-tools/cmdline-tools ${ANDROID_HOME}/cmdline-tools/latest \
    && rm /tmp/tools.zip

RUN n=0; \
    until [ "$n" -ge 4 ]; do \
        yes | sdkmanager --licenses >/dev/null \
            && sdkmanager --install "platforms;android-35" "platform-tools" "build-tools;35.0.0" \
            && break; \
        n=$((n + 1)); \
        rm -rf "${ANDROID_HOME}/.temp" "${ANDROID_HOME}/.downloadIntermediates"; sleep 10; \
    done; [ "$n" -lt 4 ]

# Gradle (AGP 8.5 needs Gradle 8.7+).
RUN curl -fsSL --retry 3 --max-time 600 "https://services.gradle.org/distributions/gradle-8.9-bin.zip" -o /tmp/gradle.zip \
    && mkdir -p /opt/gradle && unzip -q /tmp/gradle.zip -d /opt/gradle && rm /tmp/gradle.zip

ARG LD_ANDROID_SDK_VERSION=5.13.1
ARG LD_ANDROID_OBSERVABILITY_VERSION=0.60.0
ARG KOTLIN_VERSION=2.0.21
ARG AGP_VERSION=8.5.2

WORKDIR /opt/compose-lib

RUN printf '%s\n' \
    'pluginManagement {' \
    '    repositories { google(); mavenCentral(); gradlePluginPortal() }' \
    '}' \
    'dependencyResolutionManagement {' \
    '    repositories { google(); mavenCentral() }' \
    '}' \
    'rootProject.name = "compose-lib"' \
    'include(":lib")' > settings.gradle.kts \
 && printf '%s\n' \
    'plugins {' \
    "    id(\"com.android.library\") version \"${AGP_VERSION}\" apply false" \
    "    id(\"org.jetbrains.kotlin.android\") version \"${KOTLIN_VERSION}\" apply false" \
    "    id(\"org.jetbrains.kotlin.plugin.compose\") version \"${KOTLIN_VERSION}\" apply false" \
    '}' > build.gradle.kts \
 && mkdir -p lib/src/main/kotlin/com/launchdarkly/masking lib/src/main/res/layout \
 && printf '%s\n' \
    'plugins {' \
    '    id("com.android.library")' \
    '    id("org.jetbrains.kotlin.android")' \
    '    id("org.jetbrains.kotlin.plugin.compose")' \
    '}' \
    'android {' \
    '    namespace = "com.launchdarkly.masking"' \
    '    compileSdk = 35' \
    '    defaultConfig { minSdk = 24 }' \
    '    compileOptions { sourceCompatibility = JavaVersion.VERSION_1_8; targetCompatibility = JavaVersion.VERSION_1_8 }' \
    '    kotlinOptions { jvmTarget = "1.8" }' \
    '    buildFeatures { compose = true }' \
    '}' \
    'dependencies {' \
    "    implementation(\"com.launchdarkly:launchdarkly-android-client-sdk:${LD_ANDROID_SDK_VERSION}\")" \
    "    implementation(\"com.launchdarkly:launchdarkly-observability-android:${LD_ANDROID_OBSERVABILITY_VERSION}\")" \
    '    implementation(platform("androidx.compose:compose-bom:2024.09.02"))' \
    '    implementation("androidx.compose.ui:ui")' \
    '    implementation("androidx.compose.material:material")' \
    '    implementation("androidx.compose.foundation:foundation")' \
    '    implementation("androidx.compose.runtime:runtime")' \
    '    implementation("androidx.appcompat:appcompat:1.7.0")' \
    '}' > lib/build.gradle.kts \
 && printf '%s\n' '<manifest xmlns:android="http://schemas.android.com/apk/res/android" />' > lib/src/main/AndroidManifest.xml \
 && printf '%s\n' \
    '<?xml version="1.0" encoding="utf-8"?>' \
    '<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"' \
    '    android:layout_width="match_parent" android:layout_height="match_parent">' \
    '    <EditText android:id="@+id/password"' \
    '        android:layout_width="match_parent" android:layout_height="wrap_content" />' \
    '</LinearLayout>' > lib/src/main/res/layout/activity_login.xml \
 && printf '%s\n' 'package com.launchdarkly.masking' 'class Placeholder' > lib/src/main/kotlin/com/launchdarkly/masking/Placeholder.kt \
 && printf '%s\n' 'android.useAndroidX=true' 'android.enableJetifier=false' 'org.gradle.jvmargs=-Xmx2g' > gradle.properties \
 && gradle -q :lib:compileDebugKotlin --no-daemon >/dev/null 2>&1 || true

WORKDIR /work
COPY shared /harness-shared
COPY languages/android-compose/harness /harness
ENTRYPOINT ["/harness/run.sh"]
