---
id: android-client-sdk/sdk-docs/features/testdata/configure-java
sdk: android-client-sdk
kind: reference
lang: java
description: Test data source configuration for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only
---

```java
import com.launchdarkly.sdk.*;
import com.launchdarkly.sdk.android.*;
import com.launchdarkly.sdk.android.integrations.*;

TestData td = TestData.dataSource();
// You can set any initial flag states here with td.update

LDConfig config = new LDConfig.Builder(AutoEnvAttributes.Enabled)
  .mobileKey("example-mobile-key")
  .dataSource(td)
  .build();
LDClient client = LDClient.init(this.getApplication(), config, context, secondsToBlock);
```
