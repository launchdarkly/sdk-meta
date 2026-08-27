---
id: cpp-server-sdk/sdk-info/init
sdk: cpp-server-sdk
kind: init
lang: cpp
file: cpp-server-sdk/init.txt
description: |
  Client initialization snippet for cpp-server-sdk (v3 native API). A
  complete runnable main.cpp modeled on cpp-sdks'
  examples/hello-cpp-server. The wrappee owns the whole program, so the
  harness asserts the snippet's own success line via SNIPPET_SUCCESS_RE
  (no scaffold can append the EXAM-HELLO trailer around a body-defined
  main).
validation:
  runtime: cpp-server
  placeholders:
    YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
  env:
    SNIPPET_SUCCESS_RE: SDK successfully initialized
---

```cpp
#include <launchdarkly/server_side/client.hpp>
#include <launchdarkly/server_side/config/config_builder.hpp>

#include <chrono>
#include <iostream>

using namespace launchdarkly;
using namespace launchdarkly::server_side;

int main() {
    // This is your LaunchDarkly SDK key.
    // Never hardcode your SDK key in production.
    auto config = ConfigBuilder("YOUR_SDK_KEY").Build();
    if (!config) {
        std::cout << "SDK failed to initialize: invalid config: "
                  << config.error() << '\n';
        return 1;
    }

    auto client = Client(std::move(*config));

    auto start_result = client.StartAsync();
    // Wait up to five seconds for the client to connect to LaunchDarkly.
    if (auto status = start_result.wait_for(std::chrono::seconds(5));
        status == std::future_status::ready && start_result.get()) {
        std::cout << "SDK successfully initialized" << std::endl;
        return 0;
    }

    std::cout << "SDK failed to initialize" << std::endl;
    return 1;
}
```
