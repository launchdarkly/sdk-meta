FROM maven:3.9-eclipse-temurin-17

# python3 is used by the harness's import-lift pre-step (see run.sh)
# to hoist `import …;` lines out of the doc-fragment body before
# javac runs — Java forbids imports inside method bodies.
RUN apt-get update && apt-get install -y --no-install-recommends \
        python3 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /work

COPY shared /harness-shared
COPY languages/jvm/harness /harness

ENTRYPOINT ["/harness/run.sh"]
