---
id: flutter-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-automatic-custom-property-population-2-0-syntax-setting-os-and-device
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "2.0 syntax, setting os and device in section \"Understanding changes to automatic custom property population\""
---

```dart
LDContextBuilder builder = LDContextBuilder();

builder.kind('user', 'example-user-key')
    .name('Sandy');

builder.kind('device', 'example-device-key')
    .set('os', LDValue.ofString('Android 25'))
    .set('device', LDValue.ofString('Pixel XL marlin'));

LDContext context = builder.build();
```
