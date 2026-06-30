FROM maven:3.9-eclipse-temurin-17

# python3 is used by the harness's import-lift pre-step (see run.sh)
# to hoist `import …;` lines out of the doc-fragment body before
# javac runs — Java forbids imports inside method bodies.
RUN apt-get update && apt-get install -y --no-install-recommends \
        python3 \
    && rm -rf /var/lib/apt/lists/*

# Pre-bake a warm maven project: the dependency set + the compiler and
# exec plugins are downloaded into ~/.m2 and a first compile is run, so
# batch per-validate cycles run fully offline (`mvn -o`) and only recompile
# the swapped source. The old per-snippet path ran `mvn clean compile
# assembly:single` from scratch every time (re-resolving plugins/deps and
# building a fat jar); here we compile once warm and run via exec:java.
WORKDIR /opt/hello-java
RUN printf '%s\n' \
    '<?xml version="1.0" encoding="UTF-8"?>' \
    '<project xmlns="http://maven.apache.org/POM/4.0.0">' \
    '  <modelVersion>4.0.0</modelVersion>' \
    '  <groupId>com.launchdarkly.tutorial</groupId>' \
    '  <artifactId>hello-java</artifactId>' \
    '  <version>1.0-SNAPSHOT</version>' \
    '  <properties>' \
    '    <maven.compiler.source>17</maven.compiler.source>' \
    '    <maven.compiler.target>17</maven.compiler.target>' \
    '    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>' \
    '  </properties>' \
    '  <dependencies>' \
    '    <dependency><groupId>com.launchdarkly</groupId><artifactId>launchdarkly-java-server-sdk</artifactId><version>7.13.4</version></dependency>' \
    '    <dependency><groupId>com.launchdarkly</groupId><artifactId>launchdarkly-java-server-sdk-otel</artifactId><version>0.2.0</version></dependency>' \
    '    <dependency><groupId>io.opentelemetry</groupId><artifactId>opentelemetry-sdk</artifactId><version>1.40.0</version></dependency>' \
    '    <dependency><groupId>com.launchdarkly</groupId><artifactId>launchdarkly-java-server-sdk-redis-store</artifactId><version>3.1.1</version></dependency>' \
    '    <dependency><groupId>com.launchdarkly</groupId><artifactId>launchdarkly-java-server-sdk-dynamodb-store</artifactId><version>5.0.0</version></dependency>' \
    '    <dependency><groupId>com.launchdarkly</groupId><artifactId>launchdarkly-java-server-sdk-consul-store</artifactId><version>5.0.0</version></dependency>' \
    '    <dependency><groupId>dev.openfeature</groupId><artifactId>sdk</artifactId><version>1.7.0</version></dependency>' \
    '    <dependency><groupId>com.launchdarkly</groupId><artifactId>launchdarkly-openfeature-serverprovider</artifactId><version>1.0.0</version></dependency>' \
    '  </dependencies>' \
    '  <build><plugins>' \
    '    <plugin><groupId>org.codehaus.mojo</groupId><artifactId>exec-maven-plugin</artifactId><version>3.1.1</version></plugin>' \
    '  </plugins></build>' \
    '</project>' > pom.xml \
    && mkdir -p src/main/java/com/launchdarkly/tutorial \
    && printf 'package com.launchdarkly.tutorial;\npublic class App { public static void main(String[] a) { System.out.println("placeholder"); } }\n' \
        > src/main/java/com/launchdarkly/tutorial/App.java \
    && mvn -q -B compile \
    && mvn -q -B exec:java -Dexec.mainClass=com.launchdarkly.tutorial.App \
    && cp pom.xml /opt/baseline-pom.xml

WORKDIR /work

COPY shared /harness-shared
COPY languages/jvm/harness /harness

ENTRYPOINT ["/harness/run.sh"]
