---
id: java-server-sdk/getting-started/pom-build
sdk: java-server-sdk
kind: manifest-fragment
lang: xml
description: pom.xml `<build>` block configuring the maven-assembly-plugin.
ld-application:
  slot: pom-build
---

Configure the Maven Assembly Plugin in your `pom.xml` to make it easier to run the application:

```xml
<build>
  <plugins>
    <plugin>
      <artifactId>maven-assembly-plugin</artifactId>
      <configuration>
        <archive>
          <manifest>
            <mainClass>com.launchdarkly.tutorial.App</mainClass>
          </manifest>
        </archive>
        <descriptorRefs>
          <descriptorRef>jar-with-dependencies</descriptorRef>
        </descriptorRefs>
      </configuration>
    </plugin>
  </plugins>
</build>
```
