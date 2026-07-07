---
id: android-client-sdk/observability/custom-masking-compose-mask-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: lib/src/main/kotlin/com/launchdarkly/masking/Snippet.kt
description: "Android observability plugin: Custom masking with ldMask, Compose (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-masking
---

```kotlin
import com.launchdarkly.observability.api.ldMask

@Composable
fun CreditCardField() {
    var number by remember { mutableStateOf("") }
    TextField(
        value = number,
        onValueChange = { number = it },
        modifier = Modifier
            .fillMaxWidth()
            .ldMask() // mask this composable in session replay
    )
}
```
