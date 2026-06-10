---
id: roku-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: Flag evaluation reason example for Roku (BrightScript).
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only

---

```brightscript
config.setUseEvaluationReasons(true)

details = launchDarkly.intVariationDetail("example-flag-key", 123)
```
