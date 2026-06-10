---
id: cpp-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-c-sdk-v2-cpp
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Flag evaluation reason example for the C client SDK v2.x (C++ binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-cpp

---

```cpp
struct LDConfig *config = LDConfigNew("example-mobile-key");
LDConfigSetUseEvaluationReasons(config, LDBooleanTrue);
LDClientCPP *client = LDClientCPP::Init(config, user, maxwaitmilliseconds);

LDVariationDetails details;

bool value = client->boolVariationDetail("your.feature.key", false, &details);

/* inspect details.reason, which is a JSON object */
if (strcmp(LDGetText(LDObjectLookup(details.reason, "errorKind")), "FLAG_NOT_FOUND") == 0) {
/* ... */
}

LDFreeDetailContents(details);
```
