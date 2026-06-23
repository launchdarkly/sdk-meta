---
id: android-client-sdk/sdk-docs/features/monitoring/status-listener-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Status listener registration in an Activity for Android (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
class MainActivity : Activity() {
    private var client: LDClient? = null
    private var ldStatusListener: LDStatusListener? = null

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        ldStatusListener = object : LDStatusListener {
            override fun onConnectionModeChanged(connectionInfo: ConnectionInformation) {
                // handle new connection info
            }

            override fun onInternalFailure(ldFailure: LDFailure) {
                // handle failure
            }
        }

        client = LDClient.get()
        client?.registerStatusListener(ldStatusListener)
    }

    override fun onDestroy() {
        super.onDestroy()
        client?.unregisterStatusListener(ldStatusListener)
    }
}
```
