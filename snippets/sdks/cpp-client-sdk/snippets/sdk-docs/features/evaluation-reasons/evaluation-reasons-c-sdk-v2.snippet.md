---
id: cpp-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-c-sdk-v2
sdk: cpp-client-sdk
kind: reference
lang: c
description: Flag evaluation reason example for the C client SDK v2.x (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-c

---

```c
struct LDConfig *config = LDConfigNew("example-mobile-key");
LDConfigSetUseEvaluationReasons(config, LDBooleanTrue);
struct LDClient *client = LDClientInit(config, user, maxwaitmilliseconds);

LDVariationDetails details;
LDBoolean value;

value = LDBoolVariationDetail(client, "your.feature.key", LDBooleanFalse, &details);

/* inspect details.reason, which is a JSON object */
if (strcmp(LDGetText(LDObjectLookup(details.reason, "errorKind")), "FLAG_NOT_FOUND") == 0) {
    /* ... */
}

LDFreeDetailContents(details);
```
