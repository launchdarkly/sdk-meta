# Build context is `validators/`. See validators/languages/python/Dockerfile
# for the rationale.
#
# Type-checking TypeScript validator for the edge SDK reference-page doc
# fragments (akamai, cloudflare, fastly, vercel). Unlike the parse-only
# `edge-ts` validator, this one pre-installs the real edge SDK packages
# plus each edge runtime's ambient type packages and runs `tsc --noEmit`
# with full module resolution and type-checking, so a fragment that calls
# the SDK with the wrong method name or argument shape fails. Edge SDKs
# run in edge runtimes (Workers, Compute@Edge, EdgeWorkers, Vercel Edge)
# that can't execute against a live LD environment in CI, so type-checking
# against the real published types is the strongest validation available.
FROM node:20-bookworm-slim

# Pin so the validator is reproducible; bump for newer TS syntax support.
ARG TYPESCRIPT_VERSION=5.7.2

WORKDIR /opt/edge-tsc
# --legacy-peer-deps: the four edge SDKs pin overlapping ranges of the
# shared js-server-sdk-common peer at slightly different versions; we only
# need each package's published .d.ts on disk for type-checking, so a
# lenient install is fine.
RUN npm init -y >/dev/null 2>&1 \
    && npm install --no-audit --no-fund --no-progress --legacy-peer-deps \
        typescript@${TYPESCRIPT_VERSION} \
        @launchdarkly/akamai-server-edgekv-sdk \
        @launchdarkly/cloudflare-server-sdk \
        @launchdarkly/fastly-server-sdk \
        @launchdarkly/vercel-server-sdk \
        @launchdarkly/js-server-sdk-common \
        @vercel/edge-config \
        @cloudflare/workers-types \
        @fastly/js-compute \
        @types/node

WORKDIR /work
COPY shared /harness-shared
COPY languages/edge-tsc/harness /harness
ENTRYPOINT ["/harness/run.sh"]
