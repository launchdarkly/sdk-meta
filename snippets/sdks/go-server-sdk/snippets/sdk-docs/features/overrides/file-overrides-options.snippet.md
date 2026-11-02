---
id: go-server-sdk/sdk-docs/features/overrides/file-overrides-options
sdk: go-server-sdk
kind: reference
lang: go
description: File-based flag override source with multiple files, duplicate key handling, and change detection for Go SDK v7.18+.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
overrides := ldoverrides.FileSource().
    FilePaths("/etc/launchdarkly/overrides.json", "/etc/launchdarkly/team-overrides.json").
    DuplicateKeysHandling(ldoverrides.DuplicateKeysKeepFirst).
    ChangeDetection(ldoverrides.Watching)

config.DataSystem = ldcomponents.DataSystem().Default().Overrides(overrides)
```
