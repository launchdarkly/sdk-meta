FROM debian:bookworm-slim

# Parse-and-type-check validator for the legacy v2.x C++ binding of
# the cpp-client SDK (the `LDClientCPP::Init` / `client->boolVariation`
# surface exposed via <launchdarkly/api.hpp>, distinct from the bare
# C surface in cpp-client-v2-c). Removed when cpp-client was rewritten
# on top of the modern cpp-sdks codebase.
#
# Ships a stub <launchdarkly/api.hpp> declaring the v2 C++ binding's
# class + free-function surface as inline no-ops. g++ -fsyntax-only
# accepts the staged source against the stubs — no real client links.

RUN apt-get update && apt-get install -y --no-install-recommends \
        g++ libc6-dev ca-certificates \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /usr/local/include/launchdarkly
COPY languages/cpp-client-v2-cpp/api.hpp /usr/local/include/launchdarkly/api.hpp

WORKDIR /work

COPY shared /harness-shared
COPY languages/cpp-client-v2-cpp/harness /harness

ENTRYPOINT ["/harness/run.sh"]
