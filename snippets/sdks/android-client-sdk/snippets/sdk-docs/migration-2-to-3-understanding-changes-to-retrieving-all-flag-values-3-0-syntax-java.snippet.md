---
id: android-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-retrieving-all-flag-values-3-0-syntax-java
sdk: android-client-sdk
kind: reference
lang: java
description: "3.0 syntax (Java) in section \"Understanding changes to retrieving all flag values\""
---

```java
Map<String, LDValue> flagValues = client.allFlags();
for (Map.Entry<String, LDValue> flag : flagValues.entrySet()) {
    switch (flag.getValue().getType()) {
        case LDValueType.NULL: // Do something with flag missing value
            break;
        case LDValueType.BOOLEAN: // Do something with boolean flag
            break;
        case LDValueType.NUMBER: // Do something with numeric flag
            break;
        case LDValueType.STRING: // Do something with string flag
            break;
        case LDValueType.ARRAY: // Do something with array flag
            break;
        case LDValueType.OBJECT: // Do something with object flag
            break;
    }
}
```
