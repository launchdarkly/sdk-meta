---
id: android-client-sdk/sdk-docs/features/monitoring/failure-types-java
sdk: android-client-sdk
kind: reference
lang: java
description: Failure type inspection for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
LDClient client = LDClient.get();
ConnectionInformation connectionInfo = client.getConnectionInformation();
LDFailure ldFailure = connectionInfo.getLastFailure();
if (ldFailure != null) {
    Timber.d("Received failure with message %s", ldFailure.getMessage());
    // Retrieve the failure type
    LDFailure.FailureType failureType = ldFailure.getFailureType();
    switch (failureType) {
        case INVALID_RESPONSE_BODY:
            Timber.d("Received invalid response body");
            break;
        case NETWORK_FAILURE:
            Timber.d("Network failure, may have bad connection");
            break;
        case UNEXPECTED_STREAM_ELEMENT_TYPE:
            Timber.d("Unexpected stream element, may require update");
            break;
        case UNEXPECTED_RESPONSE_CODE:
            LDInvalidResponseCodeFailure responseCodeFailure =
                (LDInvalidResponseCodeFailure) ldFailure;
            int responseCode = responseCodeFailure.getResponseCode();
            if (responseCodeFailure.isRetryable()) {
                Timber.d("Received invalid response code %d", responseCode);
            } else {
                Timber.d("Received invalid response code %d, giving up", responseCode);
            }
            break;
        case UNKNOWN_ERROR:
        default:
            Timber.d("Unknown error");
            break;
    }

    Throwable cause = ldFailure.getCause();
    if (cause != null) {
        // Do something with underlying cause
    }
}
```
