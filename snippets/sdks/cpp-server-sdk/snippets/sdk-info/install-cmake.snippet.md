---
id: cpp-server-sdk/sdk-info/install-cmake
sdk: cpp-server-sdk
kind: install
lang: cmake
file: cpp-server-sdk/install-cmake.txt
description: |
  Install fragment for cpp-server-sdk (CMake). Declarative build wiring,
  not runnable shell — unvalidated, same disposition as the haskell
  cabal/stack fragments. The SDK is incorporated by cloning
  launchdarkly/cpp-sdks as a subdirectory and linking the
  launchdarkly::server target.
---

```cmake
# Clone https://github.com/launchdarkly/cpp-sdks as a subdirectory of
# your project, then add it to your CMakeLists.txt:
add_subdirectory(cpp-sdks)
target_link_libraries(your-target PRIVATE launchdarkly::server)
```
