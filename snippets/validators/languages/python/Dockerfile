# Build context is `validators/`. The Dockerfile references
# `languages/python/...` for its own files and `shared/...` for the shared
# harness library that other validators also use.
ARG PYTHON_IMAGE
FROM ${PYTHON_IMAGE}

WORKDIR /snippet

COPY shared /harness-shared
COPY languages/python/harness /harness

ENTRYPOINT ["/harness/run.sh"]
