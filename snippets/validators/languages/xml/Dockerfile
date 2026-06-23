# Build context is `validators/`. See validators/languages/python/Dockerfile
# for the rationale.
#
# Well-formedness validator for XML doc fragments (e.g. the Roku
# SceneGraph component definitions). xmllint parses the staged file and
# exits non-zero on any syntax error. Parse-only: no schema or DTD
# validation -- doc fragments reference vendor component types we don't
# ship -- and no LD env.
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
        libxml2-utils \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /work

COPY shared /harness-shared
COPY languages/xml/harness /harness

ENTRYPOINT ["/harness/run.sh"]
