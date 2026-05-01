---
id: cpp-client-sdk/scaffolds/cpp-client-syntax-only
sdk: cpp-client-sdk
kind: scaffold
lang: cpp
file: main.cpp
description: |
  Parse-only validator for C++ client SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: cpp-client
  entrypoint: main.cpp
---

```cpp
#include <iostream>

void _wrappee() {
{{ body }}
}

int main() {
    std::cout << "feature flag evaluates to true" << std::endl;
    return 0;
}
```
