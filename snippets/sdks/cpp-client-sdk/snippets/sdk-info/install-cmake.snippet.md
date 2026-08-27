---
id: cpp-client-sdk/sdk-info/install-cmake
sdk: cpp-client-sdk
kind: install
lang: cmake
file: cpp-client-sdk/install-cmake.txt
description: |
  Install fragment for cpp-client-sdk (CMake). Declarative build wiring,
  not runnable shell — unvalidated. The SDK is incorporated by cloning
  launchdarkly/cpp-sdks as a subdirectory and linking the
  launchdarkly::client target.
---

```cmake
# Clone https://github.com/launchdarkly/cpp-sdks as a subdirectory of
# your project, then add it to your CMakeLists.txt:
add_subdirectory(cpp-sdks)
target_link_libraries(your-target PRIVATE launchdarkly::client)
```
