---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-automatic-custom-property-population-4-0-syntax-setting-device-and-os
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax, setting device and os in section \"Understanding changes to automatic custom property population\""
---

```java
LDContext context = LDContext.builder("example-context-key")
    .set("os", "25")
    .set("device", "Pixel XL marlin")
    .build();
```
