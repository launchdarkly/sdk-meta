---
id: lua-server-sdk/scaffolds/lua-syntax-only
sdk: lua-server-sdk
kind: scaffold
lang: lua
file: main.lua
description: |
  Parse-only validator for Lua server SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: lua-server
  entrypoint: main.lua
---

```lua
local function _wrappee()
{{ body }}
end

print("feature flag evaluates to true")
```
