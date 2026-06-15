---
id: android-client-sdk/experimentation/full
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Full experimentation onboarding for android-client-sdk — initialize, identify off the main thread on login/eligibility, evaluate, and track conversions.
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only
---

```kotlin
import com.launchdarkly.sdk.LDContext
import com.launchdarkly.sdk.LDValue
import com.launchdarkly.sdk.android.LDConfig.Builder.AutoEnvAttributes
import com.launchdarkly.sdk.android.LDClient
import com.launchdarkly.sdk.android.LDConfig
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext

// Initialize once — extra clients can cause inconsistent experiment results.
val ldConfig: LDConfig = LDConfig.Builder(AutoEnvAttributes.Enabled)
    .mobileKey("YOUR_MOBILE_KEY")
    .build()

// A "context" is a data object representing users, devices, organizations, and other entities.
// If you already know the user's key at startup, initialize with it directly.
// Use a consistent key so the same user gets the same experience.
val context: LDContext = LDContext.create("EXAMPLE_CONTEXT_KEY")

class MainActivity : AppCompatActivity() {

    private lateinit var client: LDClient

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        client = LDClient.init(application, ldConfig, context, 0)
    }

    // Call this when the user logs in or becomes eligible mid-session.
    // Wait for identify to finish before evaluating experiment flags.
    fun onUserEligible(userKey: String) {
        // Use the logged-in user's ID so experiment assignment stays consistent.
        val finalCtx = LDContext.builder(userKey)
            .kind("user")
            // any attributes that affect targeting or eligibility
            .build()

        CoroutineScope(Dispatchers.IO).launch {
            // identify() returns a Future; await its completion off the main thread.
            client.identify(finalCtx).get()

            // Evaluate the experiment flag where the user encounters the experience.
            val variant = client.stringVariation("YOUR_FLAG_KEY", "control")

            withContext(Dispatchers.Main) {
                applyVariant(variant)
            }
        }
    }

    // Call this when the user completes a metric action.
    // Use the same user key you used when evaluating the flag — mismatched keys break conversion tracking.
    // The data argument is optional and accepts any shape your metric needs.
    fun trackMetric(metricKey: String, data: LDValue = LDValue.ofNull()) {
        client.trackData(metricKey, data)
    }

    // The SDK batches and flushes events automatically, including when the app is
    // backgrounded. Don't add manual flush() calls or onStop() flushes — they're
    // unnecessary and make real problems harder to spot.
    // Don't skip or cache flag evaluations to reduce exposure counts — LaunchDarkly deduplicates them automatically.
}
```
