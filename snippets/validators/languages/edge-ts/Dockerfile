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
ARG NODE_BOOKWORM_SLIM_IMAGE
FROM ${NODE_BOOKWORM_SLIM_IMAGE}

COPY shared/versions/npm.env /versions/npm.env
RUN . /versions/npm.env && npm install -g npm@"$NPM_VERSION" \
 && printf 'min-release-age=3\nmin-release-age-exclude[]=@launchdarkly/*\nmin-release-age-exclude[]=launchdarkly-*\n' > /root/.npmrc

# Pin so the validator is reproducible; bump for newer TS syntax support.
ARG EDGE_TYPESCRIPT_VERSION
WORKDIR /opt/edge-ts
RUN npm install --no-audit --no-fund typescript@${EDGE_TYPESCRIPT_VERSION}

WORKDIR /work
COPY shared /harness-shared
COPY languages/edge-ts/harness /harness
ENTRYPOINT ["/harness/run.sh"]
