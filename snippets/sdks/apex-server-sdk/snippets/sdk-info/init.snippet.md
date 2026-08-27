---
id: apex-server-sdk/sdk-info/init
sdk: apex-server-sdk
kind: init
lang: java
file: apex-server-sdk/init.txt
description: |
  Client initialization snippet for apex-server-sdk. Parse-checked as
  anonymous Apex via the apex-syntax-only scaffold. The SDK key never
  appears in Apex code — flag data reaches Salesforce through the bridge
  application, which reads LD_SDK_KEY.
validation:
  scaffold: apex-server-sdk/scaffolds/apex-syntax-only
---

```java
// The Apex SDK receives flag data through the bridge application,
// which reads your SDK key from the LD_SDK_KEY environment variable —
// the key never appears in Apex code.
LDConfig config = new LDConfig.Builder().build();
LDClient client = new LDClient(config);
```
