#!/bin/sh
# Runs the staged C++ client snippet against a real LaunchDarkly environment.
# Mirrors gonfalon's Get Started flow: clone cpp-sdks alongside main.cpp,
# add it via CMake, link the client SDK target. The Dockerfile pre-cloned
# cpp-sdks at /opt/cpp-sdks and prewarmed the build cache, so per-validate
# cycles only compile the user's main.cpp.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_MOBILE_KEY LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT
cd "$WORK"

cp "/snippet/$SNIPPET_ENTRYPOINT" main.cpp
ln -s /opt/cpp-sdks cpp-sdks

cat > CMakeLists.txt <<'EOF'
cmake_minimum_required(VERSION 3.19)
project(hello-cpp-client LANGUAGES CXX)
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(THREADS_PREFER_PTHREAD_FLAG ON)
find_package(Threads REQUIRED)
add_subdirectory(cpp-sdks)
add_executable(hello main.cpp)
target_link_libraries(hello PRIVATE launchdarkly::client Threads::Threads)
EOF

mkdir build
cd build
cmake -G Ninja -DBUILD_TESTING=OFF .. >/tmp/cmake.log 2>&1 \
    || { cat /tmp/cmake.log >&2; exit 1; }
cmake --build . --target hello >/tmp/build.log 2>&1 \
    || { cat /tmp/build.log >&2; exit 1; }

LOG=$(mktemp)

timeout --signal=TERM 60s ./hello >"$LOG" 2>&1 &
PID=$!

deadline=$(( $(date +%s) + 55 ))
if await_success_line "$LOG" "$PID" "$deadline"; then
    exit 0
fi

kill -TERM "$PID" 2>/dev/null || true
wait "$PID" 2>/dev/null || true
fail_with_log "$LOG" "did not see expected line: feature flag evaluates to true"
