---
id: java-server-sdk/sdk-docs/features/overrides/file-overrides-options
sdk: java-server-sdk
kind: reference
lang: java
description: File-based flag override source with multiple files, duplicate key handling, and change detection for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
FileOverrideSourceBuilder overrides = FileOverrides.source()
    .filePaths("/etc/launchdarkly/overrides.json", "/etc/launchdarkly/team-overrides.json")
    .duplicateKeysHandling(FileData.DuplicateKeysHandling.IGNORE)
    .changeDetection(FileOverrides.ChangeDetection.WATCHING);

LDConfig config = new LDConfig.Builder()
    .dataSystem(Components.dataSystem().defaultMode().overrides(overrides))
    .build();
```
