---
id: sdk-docs/features/filedata/flag-data-json
kind: reference
lang: json
description: Example full-format flag data file with flags and segments.
---

```json
{
  "flags": {
    "flag-key-1": {
      "key": "flag-key-1",
        "on": true,
        "variations": [ "a", "b" ]
      }
  },
  "segments": {
    "segment-key-1": {
      "key": "segment-key-1",
      "includes": [ "context-key-1" ]
    }
  }
}
```
