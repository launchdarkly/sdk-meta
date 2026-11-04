---
id: java-server-sdk/sdk-docs/using-the-java-sdk-in-osgi-xml
sdk: java-server-sdk
kind: reference
lang: xml
description: "XML in section \"Using the Java SDK in OSGi\""
# TODO(validate): body mixes a Maven <dependency> block and a Gradle string, so it is not well-formed XML.
# To validate the XML part with the xml runtime, move the Gradle line to a separate snippet.
---

```xml
<!-- in Maven -->
<dependency>
   <groupId>com.launchdarkly</groupId>
   <artifactId>launchdarkly-java-server-sdk</artifactId>
   <version>7.17.3</version>
   <classifier>all</classifier>
</dependency>

<!-- or in Gradle -->
"com.launchdarkly:launchdarkly-java-server-sdk:7.17.3:all"
```
