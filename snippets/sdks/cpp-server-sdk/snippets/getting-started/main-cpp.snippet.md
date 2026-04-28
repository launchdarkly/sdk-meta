---
id: cpp-server-sdk/getting-started/main-cpp
sdk: cpp-server-sdk
kind: hello-world
lang: cpp
file: main.cpp
description: Hello-world program that initializes the C++ server SDK and evaluates a feature flag.
inputs:
  apiKey:
    type: sdk-key
    description: SDK key baked into the rendered source.
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source.
ld-application:
  slot: main-cpp
# Validator pending. Per-validate cycle requires a Docker image with
# cmake + boost + openssl + ninja and a checkout of cpp-sdks; first
# build is multi-minute even with prebuilt deps.
---

Create a file named `main.cpp` add the following code:

```cpp
#include <launchdarkly/context_builder.hpp>
#include <launchdarkly/server_side/client.hpp>
#include <launchdarkly/server_side/config/config_builder.hpp>

#include <cstring>
#include <iostream>

// Set INIT_TIMEOUT_MILLISECONDS to the amount of time you will wait for
// the client to become initialized.
#define INIT_TIMEOUT_MILLISECONDS 3000

using namespace launchdarkly;
using namespace launchdarkly::server_side;

int main() {
    auto config = ConfigBuilder("{{ apiKey }}").Build();
    if (!config) {
        std::cout << "error: config is invalid: " << config.error() << std::endl;
        return 1;
    }

    auto client = Client(std::move(*config));

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

    auto const context =
        ContextBuilder().Kind("user", "example-user-key").Name("Sandy").Build();

    bool const flag_value =
        client.BoolVariation(context, "{{ featureKey }}", false);

    std::cout << "*** Feature flag '{{ featureKey }}' is "
              << (flag_value ? "true" : "false") << std::endl;

    return 0;
}
```
