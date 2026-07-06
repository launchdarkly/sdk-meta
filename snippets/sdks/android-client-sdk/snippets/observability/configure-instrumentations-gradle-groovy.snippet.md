---
id: android-client-sdk/observability/configure-instrumentations-gradle-groovy
sdk: android-client-sdk
kind: reference
lang: groovy
file: android-client-sdk/observability/configure-instrumentations-gradle-groovy.gradle
description: "Android observability plugin in section \"Configure additional instrumentations\" (Gradle Groovy)"
# Not validated: a build.gradle plugins+dependencies configuration block, not compilable app code.
---

```java
plugins {
    id 'net.bytebuddy.byte-buddy-gradle-plugin' version '1.+'
}

dependencies {
    // Android HTTP Url instrumentation
    implementation 'io.opentelemetry.android.instrumentation:httpurlconnection-library:0.11.0-alpha'
    byteBuddy 'io.opentelemetry.android.instrumentation:httpurlconnection-agent:0.11.0-alpha'

    // OkHTTP instrumentation
    implementation 'io.opentelemetry.android.instrumentation:okhttp3-library:0.11.0-alpha'
    byteBuddy 'io.opentelemetry.android.instrumentation:okhttp3-agent:0.11.0-alpha'
}
```
