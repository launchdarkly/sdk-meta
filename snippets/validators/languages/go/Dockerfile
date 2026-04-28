# Build context is `validators/`. See validators/languages/python/Dockerfile
# for the rationale.
FROM golang:1.24

WORKDIR /work

COPY shared /harness-shared
COPY languages/go/harness /harness

ENTRYPOINT ["/harness/run.sh"]
