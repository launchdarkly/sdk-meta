---
id: android-client-sdk/observability/custom-masking-view-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
file: lib/src/main/kotlin/com/launchdarkly/masking/Snippet.kt
description: "Android observability plugin: Custom masking with ldMask, View (Kotlin)"
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-masking
---

```kotlin
import com.launchdarkly.observability.api.ldMask

class LoginActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_login)

        val password = findViewById<EditText>(R.id.password)
        password.ldMask() // mask this field in session replay
    }
}
```
