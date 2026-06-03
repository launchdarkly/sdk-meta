FROM debian:bookworm-slim

# Parse-and-type-check validator for the legacy v2.x C server SDK
# (launchdarkly/c-server-sdk). The v2 C API surface (LDUserNew,
# LDBoolVariation, LDConfigNew, ...) was removed when the SDK was
# rewritten on top of the modern cpp-sdks C++ codebase. The current
# cpp-server validator builds cpp-sdks v3+, so it can't compile v2-era
# fragments.
#
# Rather than clone and build the entire c-server-sdk repo, we ship a
# minimal stub header (api.h) that declares the v2 API surface as
# inline no-ops. The body compiles against the stubs — parse +
# type-check are the goal here; no real flag evaluation runs.

RUN apt-get update && apt-get install -y --no-install-recommends \
        gcc libc6-dev ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Stub `launchdarkly/api.h` — the v2 C SDK's public header.
# Declares (but does not link) the types and functions the v2.x
# fragments reference. Inline no-op definitions so the staged
# snippet links cleanly with a single `gcc` invocation against
# `-include launchdarkly/api.h`.
RUN mkdir -p /usr/local/include/launchdarkly
COPY languages/cpp-server-v2-c/api.h /usr/local/include/launchdarkly/api.h

WORKDIR /work

COPY shared /harness-shared
COPY languages/cpp-server-v2-c/harness /harness

ENTRYPOINT ["/harness/run.sh"]
