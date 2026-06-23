---
id: android-client-sdk/sdk-docs/features/monitoring/failure-types-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Failure type inspection for Android (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
val client: LDClient = LDClient.get()
val connectionInfo: ConnectionInformation = client.connectionInformation
val ldFailure: LDFailure? = connectionInfo.lastFailure
if (ldFailure != null) {
    Timber.d("Received failure with message %s", ldFailure.message)
    // Retrieve the failure type
    val failureType: LDFailure.FailureType = ldFailure.failureType
    when (failureType) {
        INVALID_RESPONSE_BODY -> {
            Timber.d("Received invalid response body")
        }
        NETWORK_FAILURE -> {
            Timber.d("Network failure, may have bad connection")
        }
        UNEXPECTED_STREAM_ELEMENT_TYPE -> {
            Timber.d("Unexpected stream element, may require update")
        }
        UNEXPECTED_RESPONSE_CODE -> {
            val responseCodeFailure: LDInvalidResponseCodeFailure =
                ldFailure as LDInvalidResponseCodeFailure
            val responseCode = responseCodeFailure.responseCode
            if (responseCodeFailure.isRetryable) {
                Timber.d("Received invalid response code %d", responseCode)
            } else {
                Timber.d("Received invalid response code %d, giving up", responseCode)
            }
        }
        UNKNOWN_ERROR -> {
            Timber.d("Unknown error")
        }
    }
    val cause: Throwable? = ldFailure.cause
    if (cause != null) {
        // Do something with underlying cause
    }
}
```
