---
id: cpp-client-sdk/sdk-info/init
sdk: cpp-client-sdk
kind: init
lang: cpp
file: cpp-client-sdk/init.txt
description: |
  Client initialization snippet for cpp-client-sdk (v3 native API). A
  complete runnable main.cpp modeled on cpp-sdks'
  examples/hello-cpp-client. The wrappee owns the whole program, so the
  harness asserts the snippet's own success line via SNIPPET_SUCCESS_RE.
validation:
  runtime: cpp-client
  placeholders:
    YOUR_MOBILE_KEY: LAUNCHDARKLY_MOBILE_KEY
  env:
    SNIPPET_SUCCESS_RE: SDK successfully initialized
---

```cpp
#include <launchdarkly/client_side/client.hpp>
#include <launchdarkly/context_builder.hpp>

#include <chrono>
#include <iostream>

using namespace launchdarkly;
using namespace launchdarkly::client_side;

int main() {
    // This is your LaunchDarkly mobile key.
    // Mobile keys are not secret, but never embed a server-side SDK key
    // in a client-side application.
    auto config = ConfigBuilder("YOUR_MOBILE_KEY").Build();
    if (!config) {
        std::cout << "SDK failed to initialize: invalid config: "
                  << config.error() << '\n';
        return 1;
    }

    auto context =
        ContextBuilder().Kind("user", "example-user-key").Build();
    auto client = Client(std::move(*config), std::move(context));

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
