---
id: flutter-client-sdk/sdk-docs/migration-1-to-2-referencing-properties-of-an-attribute-object-2-0-syntax-context-with-object-attributes
sdk: flutter-client-sdk
kind: reference
lang: dart
description: "2.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
---

```dart
LDValue addressData = LDValueObjectBuilder()
    .addValue('street', LDValue.ofString('Main St'))
    .addValue('city', LDValue.ofString('Springfield'))
    .build();

LDContextBuilder builder = LDContextBuilder();
builder.kind('user', 'example-user-key')
    .set('address', addressData);

LDContext context = builder.build();
```
