---
id: cpp-client-sdk/getting-started/main-cpp
sdk: cpp-client-sdk
kind: hello-world
lang: cpp
file: main.cpp
description: Hello-world program that initializes the C++ client SDK and evaluates a feature flag.
inputs:
  mobileKey:
    type: mobile-key
    description: Mobile key baked into the rendered source.
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source.
ld-application:
  slot: main-cpp
# Validator pending — same toolchain story as cpp-server-sdk.
---

Create a file named `main.cpp` add the following code:

```cpp
#include <launchdarkly/client_side/client.hpp>
#include <launchdarkly/context_builder.hpp>

#include <cstring>
#include <iostream>

// Set INIT_TIMEOUT_MILLISECONDS to the amount of time you will wait for
// the client to become initialized.
#define INIT_TIMEOUT_MILLISECONDS 3000

using namespace launchdarkly;
using namespace launchdarkly::client_side;

int main() {

    auto config = ConfigBuilder("{{ mobileKey }}").Build();
    if (!config) {
        std::cout << "error: config is invalid: " << config.error() << std::endl;
        return 1;
    }

    auto context =
        ContextBuilder().Kind("user", "example-user-key").Name("Sandy").Build();

    auto client = Client(std::move(*config), std::move(context));

    auto start_result = client.StartAsync();

    if (auto const status = start_result.wait_for(
            std::chrono::milliseconds(INIT_TIMEOUT_MILLISECONDS));
        status == std::future_status::ready) {
        if (start_result.get()) {
            std::cout << "*** SDK successfully initialized!" << std::endl;
        } else {
            std::cout << "*** SDK failed to initialize" << std::endl;
            return 1;
        }
    } else {
        std::cout << "*** SDK initialization didn't complete in "
                  << INIT_TIMEOUT_MILLISECONDS << "ms" << std::endl;
        return 1;
    }

    bool const flag_value = client.BoolVariation("{{ featureKey }}", false);

    std::cout << "*** Feature flag '{{ featureKey }}' is "
              << (flag_value ? "true" : "false") << std::endl;

    return 0;
}
```
