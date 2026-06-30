---
id: java-server-sdk/sdk-docs/openfeature/initialize-the-provider
sdk: java-server-sdk
kind: reference
lang: java
file: src/main/java/com/launchdarkly/tutorial/Main.java
description: "Java OpenFeature provider in section \"Initialize the provider\""
validation:
  scaffold: java-server-sdk/scaffolds/openfeature-jvm-init-runner
  placeholders:
    YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
---

```java
public class Main {
    public static void main(String[] args) {
        OpenFeatureAPI.getInstance().setProvider(new Provider("YOUR_SDK_KEY"));

        Client client = OpenFeatureAPI.getInstance().getClient();
    }
}
```
