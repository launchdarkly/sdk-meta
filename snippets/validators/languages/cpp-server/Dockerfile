FROM ubuntu:24.04

# Build chain for the LD C++ Server SDK (cmake, ninja, boost, openssl,
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
# into the project dir at build time rather than committing the tree.
ARG CPP_SDKS_REF=launchdarkly-cpp-server-v3.10.1
RUN git clone --depth 1 --branch "${CPP_SDKS_REF}" \
        https://github.com/launchdarkly/cpp-sdks.git /opt/cpp-sdks

# Pre-bake CONFIGURED projects so batch-mode per-validate cycles skip the
# expensive CMake configure (re-processing add_subdirectory(cpp-sdks))
# entirely: after swapping main.cpp the harness only runs an incremental
# `cmake --build --target hello`, which recompiles one translation unit and
# relinks against the already-built SDK objects. Configuring once per job
# instead of once per snippet is the whole point of batch mode.
#
# Two project dirs cover the build configs the snippets need: the default
# server link, and a redis-enabled variant for the redis-integration
# fragments. Both share one CMakeLists that links the redis source library
# only when LD_BUILD_REDIS_SUPPORT is set, so the variant is selected purely
# by the configure-time -D flag.
RUN for proj in /opt/hello-cpp /opt/hello-cpp-redis; do \
        mkdir -p "$proj" && cd "$proj" && ln -s /opt/cpp-sdks cpp-sdks && \
        printf '%s\n' \
            'cmake_minimum_required(VERSION 3.19)' \
            'project(hello-cpp-server LANGUAGES CXX)' \
            'set(CMAKE_CXX_STANDARD 17)' \
            'set(CMAKE_CXX_STANDARD_REQUIRED ON)' \
            'set(THREADS_PREFER_PTHREAD_FLAG ON)' \
            'find_package(Threads REQUIRED)' \
            'add_subdirectory(cpp-sdks)' \
            'add_executable(hello main.cpp)' \
            'set(LD_EXTRA_LIBS "")' \
            'if(LD_BUILD_REDIS_SUPPORT)' \
            '  set(LD_EXTRA_LIBS launchdarkly::server_redis_source)' \
            'endif()' \
            'target_link_libraries(hello PRIVATE launchdarkly::server ${LD_EXTRA_LIBS} Threads::Threads)' \
            > CMakeLists.txt && \
        printf 'int main() { return 0; }\n' > main.cpp; \
    done && \
    cmake -G Ninja -DBUILD_TESTING=OFF -B /opt/hello-cpp/build /opt/hello-cpp && \
    cmake --build /opt/hello-cpp/build --target hello && \
    cmake -G Ninja -DBUILD_TESTING=OFF -DLD_BUILD_REDIS_SUPPORT=ON -B /opt/hello-cpp-redis/build /opt/hello-cpp-redis && \
    cmake --build /opt/hello-cpp-redis/build --target hello

WORKDIR /work

COPY shared /harness-shared
COPY languages/cpp-server/harness /harness

ENTRYPOINT ["/harness/run.sh"]
