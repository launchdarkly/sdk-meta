---
id: cpp-server-sdk/scaffolds/cpp-syntax-only
sdk: cpp-server-sdk
kind: scaffold
lang: cpp
file: main.cpp
description: |
  Parse-only validator for C++ server SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: cpp-server
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
