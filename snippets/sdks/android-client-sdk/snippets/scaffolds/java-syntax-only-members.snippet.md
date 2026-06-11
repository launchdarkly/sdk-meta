---
id: android-client-sdk/scaffolds/java-syntax-only-members
sdk: android-client-sdk
kind: scaffold
lang: java
file: app/src/main/java/com/launchdarkly/hello_android/Snippet.java
description: |
  Class-member-scope sibling of `java-syntax-only`. That scaffold
  splices the body inside `onCreate()`, which breaks for fragments
  that are themselves method declarations (Java has no local
  methods). This variant splices the body at class scope instead.

  Same `android-client` validator container in `SNIPPET_CHECK=parse`
  mode, so the body compiles against the real
  `launchdarkly-android-client-sdk` aar (plus Timber, which the
  validator's pre-baked gradle project includes for log-statement
  fragments).

  No `public` modifier on the class and a class name that differs
  from the staged file name — same package-private trick as the
  sibling scaffold, so the harness's `Snippet.java` staging glob
  picks the file up.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, spliced at class scope.
validation:
  runtime: android-client
  entrypoint: app/src/main/java/com/launchdarkly/hello_android/Snippet.java
  env:
    SNIPPET_CHECK: parse
---

```java
package com.launchdarkly.hello_android;

import com.launchdarkly.sdk.*;
import com.launchdarkly.sdk.android.*;
import timber.log.Timber;

@SuppressWarnings("unused")
class SnippetMembers {
{{ body }}
}
```
