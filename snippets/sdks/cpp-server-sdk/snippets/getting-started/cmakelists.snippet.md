---
id: cpp-server-sdk/getting-started/cmakelists
sdk: cpp-server-sdk
kind: manifest
lang: cpp
file: CMakeLists.txt
description: CMake configuration file for the hello-cpp-server project.
ld-application:
  slot: cmakelists
---

Create a `CMakeLists.txt` file with the following content:

```cpp
cmake_minimum_required(VERSION 3.19)

project(
  CPPServerQuickstart
  VERSION 0.1
  DESCRIPTION "LaunchDarkly CPP Server-side SDK Quickstart"
  LANGUAGES CXX
)

set(THREADS_PREFER_PTHREAD_FLAG ON)
find_package(Threads REQUIRED)

add_subdirectory(cpp-sdks)

add_executable(cpp-server-quickstart main.cpp)

target_link_libraries(cpp-server-quickstart
      PRIVATE
        launchdarkly::server
        Threads::Threads
)
        
```
