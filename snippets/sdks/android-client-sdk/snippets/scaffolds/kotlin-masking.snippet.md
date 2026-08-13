---
id: android-client-sdk/scaffolds/kotlin-masking
sdk: android-client-sdk
kind: scaffold
lang: kotlin
file: lib/src/main/kotlin/com/launchdarkly/masking/Snippet.kt
description: |
  Type-check validator for the observability plugin's custom-masking
  fragments (Jetpack Compose `Modifier.ldMask()`/`ldUnmask()` and the
  View `EditText.ldMask()`). Routes through the dedicated `android-compose`
  validator (Kotlin 2.0 + Compose 1.7.x) because the plugin's Compose
  masking API carries Kotlin-2.0 metadata the Kotlin-1.8 android-client
  validator can't read. The fragment is a self-contained top-level
  declaration (a `@Composable fun` or an `Activity` subclass) spliced at
  file scope, right after the imports the fragments assume; the
  fragment's own `import com.launchdarkly.observability.api.ldMask`
  stays in the import region. The View fragment's `R.layout.activity_login`
  / `R.id.password` resolve against a layout baked into the validator.
inputs:
  body:
    type: string
    description: The wrappee fragment, a top-level Compose function or Activity subclass.
validation:
  runtime: android-compose
  entrypoint: lib/src/main/kotlin/com/launchdarkly/masking/Snippet.kt
---

```kotlin
package com.launchdarkly.masking

import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import androidx.compose.material.TextField
import androidx.compose.ui.Modifier
import androidx.compose.foundation.layout.fillMaxWidth
import android.os.Bundle
import android.widget.EditText
import androidx.appcompat.app.AppCompatActivity

{{ body }}
```
