# Build context is `validators/`. See validators/languages/python/Dockerfile
# for the rationale.
#
# Parse-only TypeScript validator for the edge SDK config doc fragments
# (cloudflare, vercel, akamai-edgekv, fastly), none of which had a
# validator. Uses the TypeScript compiler's transpileModule, which does
# a syntax check + type-strip with NO module resolution and NO
# type-checking. That lets fragments import edge-only packages and
# reference ambient globals (env, process) and parse cleanly without
# those packages installed — exactly the parse-only guarantee the other
# syntax-only validators give for their languages.
FROM node:20-bookworm-slim

# Pin so the validator is reproducible; bump for newer TS syntax support.
ARG TYPESCRIPT_VERSION=5.7.2
WORKDIR /opt/edge-ts
RUN npm install --no-audit --no-fund typescript@${TYPESCRIPT_VERSION}

WORKDIR /work
COPY shared /harness-shared
COPY languages/edge-ts/harness /harness
ENTRYPOINT ["/harness/run.sh"]
