---
id: java-server-sdk/sdk-docs/features/migrations/stage-switch
sdk: java-server-sdk
kind: reference
lang: java
description: Per-stage migration structure for Java SDK v7.0.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
// define the combination of reads and writes from the new and old systems
// that should occur at each migration stage

switch (migrationVariation.getStage()) {
  case OFF:
  case DUAL_WRITE:
  case SHADOW:
  case LIVE:
  case RAMP_DOWN:
  case COMPLETE:
  default: {
    // throw an error
  }
}
```
