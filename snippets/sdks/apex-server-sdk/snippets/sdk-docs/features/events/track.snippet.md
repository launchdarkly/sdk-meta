---
id: apex-server-sdk/sdk-docs/features/events/track
sdk: apex-server-sdk
kind: reference
lang: java
description: Custom event tracking example for Apex.
validation:
  scaffold: apex-server-sdk/scaffolds/apex-syntax-only

---

```java
client.track(user, 'example-event-key', 52.3, LDValue.of('my value'));
```
