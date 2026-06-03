FROM debian:bookworm-slim

# Parse-and-type-check validator for the legacy v2.x C client SDK
# (launchdarkly/c-client-sdk — a separate library from c-server-sdk,
# with a similar but distinct public surface). Removed when the
# cpp-client SDK was rewritten on top of the modern cpp-sdks C++
# codebase.
#
# Ships a stub <launchdarkly/api.h> declaring the cpp-client v2 API
# surface (LDClientInit(config, user, maxwait), LDBoolVariation(client,
# flagKey, fallback), ...) as inline no-ops. Distinct from the
# cpp-server-v2-c validator because the client- and server-side v2 C
# SDKs shared header names but had different function signatures.

RUN apt-get update && apt-get install -y --no-install-recommends \
        gcc libc6-dev ca-certificates \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /usr/local/include/launchdarkly
COPY languages/cpp-client-v2-c/api.h /usr/local/include/launchdarkly/api.h

WORKDIR /work

COPY shared /harness-shared
COPY languages/cpp-client-v2-c/harness /harness

ENTRYPOINT ["/harness/run.sh"]
