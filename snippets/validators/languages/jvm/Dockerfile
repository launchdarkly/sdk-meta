FROM maven:3.9-eclipse-temurin-17

WORKDIR /work

COPY shared /harness-shared
COPY languages/jvm/harness /harness

ENTRYPOINT ["/harness/run.sh"]
