---
id: sdk-docs/features/overrides/override-file-json
kind: reference
lang: json
description: Example flag override file with a flagValues entry and a full flag definition (JSON).
---

```json
{
  "flagValues": {
    "my-boolean-flag-key": false,
    "my-string-flag-key": "value-1"
  },
  "flags": {
    "my-targeted-flag-key": {
      "key": "my-targeted-flag-key",
      "on": true,
      "variations": [ true, false ],
      "targets": [
        { "values": [ "context-key-123abc" ], "variation": 0 }
      ],
      "fallthrough": { "variation": 1 },
      "offVariation": 1
    }
  }
}
```
