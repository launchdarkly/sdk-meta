---
id: android-client-sdk/observability/custom-masking-compose-unmask-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: lib/src/main/kotlin/com/launchdarkly/masking/Snippet.kt
description: "Android observability plugin: Custom masking with ldUnmask, Compose (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-masking
---

```kotlin
import com.launchdarkly.observability.api.ldUnmask

@Composable
fun PublicInfoField() {
    var info by remember { mutableStateOf("") }
    TextField(
        value = info,
        onValueChange = { info = it },
        modifier = Modifier
            .fillMaxWidth()
            .ldUnmask() // explicitly unmask this field
    )
}
```
