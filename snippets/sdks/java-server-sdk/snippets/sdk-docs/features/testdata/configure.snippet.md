---
id: java-server-sdk/sdk-docs/features/testdata/configure
sdk: java-server-sdk
kind: reference
lang: java
description: Test data source configuration for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
import com.launchdarkly.sdk.*;
import com.launchdarkly.sdk.server.*;
import com.launchdarkly.sdk.server.integrations.*;

TestData td = TestData.dataSource();
// You can set any initial flag states here with td.update

LDConfig config = new LDConfig.Builder()
    .dataSource(td)
    .build();
LDClient client = new LDClient(sdkKey, config);
```
