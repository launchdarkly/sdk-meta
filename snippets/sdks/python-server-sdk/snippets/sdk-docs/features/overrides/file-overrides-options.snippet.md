---
id: python-server-sdk/sdk-docs/features/overrides/file-overrides-options
sdk: python-server-sdk
kind: reference
lang: python
description: File-based flag override source with multiple files, duplicate key handling, and change detection for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
overrides = (
    FileOverrideSourceBuilder([
        "/etc/launchdarkly/overrides.json",
        "/etc/launchdarkly/team-overrides.json",
    ])
    .duplicate_keys_handling("ignore")
    .change_detection("watching")
)

config = Config(
    "YOUR_SDK_KEY",
    datasystem_config=datasystem.default().overrides(overrides).build(),
)
```
