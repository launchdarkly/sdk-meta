FROM debian:bookworm-slim

# Parse-and-type-check validator for C++-flavored fragments of the
# legacy v2.x C server SDK (launchdarkly/c-server-sdk). The v2 server
# SDK had no separate C++ binding class — its docs simply call the C
# API from C++ (`LDDetails details; bool value = LDBoolVariation(...)`),
# so those fragments need g++ rather than the cpp-server-v2-c
# validator's C11 gcc pass (`bool`, unqualified struct names, and other
# C++-isms reject under -std=c11).
#
# Reuses the cpp-server-v2-c stub <launchdarkly/api.h> — same API
# surface, compiled as C++. Parse + type-check are the goal here; no
# real flag evaluation runs.

RUN apt-get update && apt-get install -y --no-install-recommends \
        g++ libc6-dev ca-certificates \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /usr/local/include/launchdarkly
COPY languages/cpp-server-v2-c/api.h /usr/local/include/launchdarkly/api.h

WORKDIR /work

COPY shared /harness-shared
COPY languages/cpp-server-v2-cpp/harness /harness

ENTRYPOINT ["/harness/run.sh"]
