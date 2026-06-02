---
id: android-client-sdk/scaffolds/java-syntax-only
sdk: android-client-sdk
kind: scaffold
lang: java
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.java
description: |
  Parse-and-type-check validator for Android Java doc fragments.

  Routes through the `android-client` validator container in
  `SNIPPET_CHECK=parse` mode: the harness stages this file into the
  pre-baked hello-android gradle project alongside the existing
  MainApplication / MainActivity sources, then runs
  `./gradlew compileDebugJavaWithJavac` (paired with
  compileDebugKotlin so a project that contains both source kinds
  still validates fully). A passing compile means the body parses
  AND type-checks against the real
  `com.launchdarkly:launchdarkly-android-client-sdk` aar from
  Google's Maven plus the AndroidX runtime — neither of which the
  `jvm` validator's Java + Maven path can resolve.

  The wrappee body is spliced inside `BaseApplication.onCreate()`'s
  unreachable `if (false)` block so unresolved caller surfaces
  (Activity host lifecycle, etc.) have somewhere legal to land.
  Bodies that declare local helper variables, call methods on
  `this.getApplication()`, or reference Android Application context
  resolve through the enclosing BaseApplication instance.

  Bodies with `import …` lines from inside the body block are
  handled by the harness's existing Python import-lift pre-step:
  Java only allows imports between `package` and the first
  top-level declaration, so they're hoisted to file scope before
  javac runs. The lifter applies to `.kt` and `.java` files
  identically — Java's import syntax mirrors Kotlin's closely
  enough that the existing regex matches both.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: android-client
  entrypoint: app/src/main/java/com/launchdarkly/hello_android/Snippet.java
  env:
    SNIPPET_CHECK: parse
---

```java
package com.launchdarkly.hello_android;

import android.app.Application;
import android.os.Bundle;
import androidx.appcompat.app.AppCompatActivity;
import com.launchdarkly.sdk.android.LDClient;

// No `public` modifier: Java requires public top-level classes to
// live in a file matching the class name. We need this scaffold's
// staged file to be Snippet.java (not BaseApplication.java) so the
// android-client harness's *.java staging glob picks it up, and so
// it doesn't collide with the pre-baked MainApplication.kt's
// declared application name. Package-private visibility is fine —
// the class is never instantiated.
@SuppressWarnings({"unused", "ConstantConditions"})
class BaseApplication extends Application {
    // Instance-field stubs so bodies like
    // `client.boolVariation(flagKey, true)` resolve at javac time.
    // Kotlin bodies pick these up via top-level decls in the
    // pre-baked MainApplication.kt; Java has no equivalent file-scope
    // mechanism, so we declare them as instance fields and rely on
    // `onCreate()` being an instance method.
    LDClient client;
    String flagKey;

    @Override
    public void onCreate() {
        super.onCreate();
        if (false) {
{{ body }}
        }
    }
}
```
