---
id: java-server-sdk/getting-started/pom-compiler
sdk: java-server-sdk
kind: manifest-fragment
lang: xml
description: pom.xml maven-compiler source/target levels.
ld-application:
  slot: pom-compiler
---

Depending on your Java version, you may need to change the compilation source and target level in `pom.xml`:

```xml
<maven.compiler.source>1.8</maven.compiler.source>
        <maven.compiler.target>1.8</maven.compiler.target>
```
