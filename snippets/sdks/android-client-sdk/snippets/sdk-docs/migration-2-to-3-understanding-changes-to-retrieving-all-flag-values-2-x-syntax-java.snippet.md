---
id: android-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-retrieving-all-flag-values-2-x-syntax-java
sdk: android-client-sdk
kind: reference
lang: java
description: "2.x syntax (Java) in section \"Understanding changes to retrieving all flag values\""
---

```java
Map<String, ?> flagValues = client.allFlags();
for (Map.Entry<String, ?> flag : flagValues.entrySet()) {
    if (flag.getValue() instanceOf Boolean) {
        // Do something with boolean flag
    } else if (flag.getValue() instanceOf Float) {
        // Do something with numeric flag
    } else if (flag.getValue() instanceOf String) {
        // Do something with string (or serialized Json) flag
    }
}
```
