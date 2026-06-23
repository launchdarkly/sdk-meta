---
id: sdk-docs/features/identify/multi-context
kind: reference
lang: json
description: Example multi-context payload shared by all SDKs on the identify feature page.
---

```json
{
  "kind": "multi",
  "user": {
    "key": "example-user-key",
    "name": "Sandy",
    "email": "sandy@example.com"
  },
  "device": {
    "key": "example-device-key",
    "type": "iPhone",
    "deviceId": 12345
  }
}
```
