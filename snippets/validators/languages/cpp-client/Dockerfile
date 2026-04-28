FROM ubuntu:24.04

# Build chain for the LD C++ Client SDK (cmake, ninja, boost, openssl,
# git for cloning cpp-sdks, ccache to keep per-validate rebuilds cheap).
RUN apt-get update && apt-get install -y --no-install-recommends \
        build-essential cmake ninja-build pkg-config \
        libboost-all-dev libssl-dev \
        git ca-certificates ccache \
    && rm -rf /var/lib/apt/lists/*

ENV CCACHE_DIR=/root/.ccache
ENV PATH="/usr/lib/ccache:${PATH}"

# Pre-clone cpp-sdks. The snippet's CMakeLists references this via
# `add_subdirectory(cpp-sdks)` from the project root, so we symlink it
# into the staging dir at validate time rather than committing the
# tree to /work directly.
ARG CPP_SDKS_REF=launchdarkly-cpp-client-v3.11.1
RUN git clone --depth 1 --branch "${CPP_SDKS_REF}" \
        https://github.com/launchdarkly/cpp-sdks.git /opt/cpp-sdks

# Warm the ccache by configuring + building the client SDK once. Subsequent
# per-snippet builds compile only the user's main.cpp and link against the
# already-cached objects.
RUN mkdir -p /tmp/prewarm && cd /tmp/prewarm \
    && ln -s /opt/cpp-sdks cpp-sdks \
    && printf 'cmake_minimum_required(VERSION 3.19)\n\
project(prewarm LANGUAGES CXX)\n\
set(THREADS_PREFER_PTHREAD_FLAG ON)\n\
find_package(Threads REQUIRED)\n\
add_subdirectory(cpp-sdks)\n' > CMakeLists.txt \
    && mkdir build && cd build \
    && cmake -G Ninja -DBUILD_TESTING=OFF .. \
    && cmake --build . --target launchdarkly-cpp-client \
    && cd / && rm -rf /tmp/prewarm

WORKDIR /work

COPY shared /harness-shared
COPY languages/cpp-client/harness /harness

ENTRYPOINT ["/harness/run.sh"]
