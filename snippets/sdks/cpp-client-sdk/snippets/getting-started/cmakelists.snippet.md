---
id: cpp-client-sdk/getting-started/cmakelists
sdk: cpp-client-sdk
kind: manifest
lang: text
file: CMakeLists.txt
description: CMake configuration for the hello-cpp-client project.
ld-application:
  slot: cmakelists
---

Create a `CMakeLists.txt` file with the following content:

```text
cmake_minimum_required(VERSION 3.19)

project(
  CPPClientQuickstart
  VERSION 0.1
  DESCRIPTION "LaunchDarkly CPP Client-side SDK Quickstart"
  LANGUAGES CXX
)

set(THREADS_PREFER_PTHREAD_FLAG ON)
find_package(Threads REQUIRED)

add_subdirectory(cpp-sdks)

add_executable(cpp-client-quickstart main.cpp)

target_link_libraries(cpp-client-quickstart
      PRIVATE
        launchdarkly::client
        Threads::Threads
)
        
```
