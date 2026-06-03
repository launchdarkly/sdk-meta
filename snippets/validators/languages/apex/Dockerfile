# Build context is `validators/`. See validators/languages/python/Dockerfile
# for the rationale.
#
# Parse-only validator for Apex server SDK snippets. Compiling Apex for
# real needs the Salesforce CLI plus a scratch org (the apex-server-sdk
# repo's CI does exactly that), which doesn't fit the Docker validator
# mold. But that same repo's CI also runs `prettier --check "**/*.cls"`
# via `prettier-plugin-apex`, whose `apex-anonymous` parser checks Apex
# syntax offline. We reuse that parser here to syntax-check doc
# fragments without any Salesforce credentials.
#
# prettier-plugin-apex ships its parser as a bundled JVM tool
# (apex-ast-serializer), so the image needs both Node and a JRE.
FROM node:20-bookworm-slim

# prettier-plugin-apex's apex-ast-serializer requires Java 11+. The
# bookworm default JRE is 17, which it runs under cleanly.
RUN apt-get update \
    && apt-get install -y --no-install-recommends default-jre-headless \
    && rm -rf /var/lib/apt/lists/*

# Pin both so the validator is reproducible; bump when we want newer
# Apex language-feature support. prettier 3.x is the line
# prettier-plugin-apex 2.x targets.
ARG PRETTIER_VERSION=3.4.2
ARG PRETTIER_PLUGIN_APEX_VERSION=2.2.6

# Pre-bake the install under /opt/apex so per-validate cycles skip the
# npm fetch. The harness invokes prettier from here.
WORKDIR /opt/apex
RUN npm install --no-audit --no-fund \
    prettier@${PRETTIER_VERSION} \
    prettier-plugin-apex@${PRETTIER_PLUGIN_APEX_VERSION}

# Warm the apex-ast-serializer JVM/classpath once at build time so the
# first real parse doesn't pay the cold-start unpack cost.
RUN printf 'Integer warm = 1;\n' > /tmp/warm.apex \
    && ./node_modules/.bin/prettier \
        --plugin=prettier-plugin-apex \
        --parser=apex-anonymous /tmp/warm.apex > /dev/null \
    && rm -f /tmp/warm.apex

WORKDIR /work

COPY shared /harness-shared
COPY languages/apex/harness /harness

ENTRYPOINT ["/harness/run.sh"]
