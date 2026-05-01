---
id: lua-server-sdk/sdk-docs/evaluate-a-context-lua-sdk-v1-x
sdk: lua-server-sdk
kind: reference
lang: lua
description: "Lua SDK v1.x in section \"Evaluate a context\""
---

```lua
if client:boolVariation(context, "example-flag-key", false) then
    print "feature is enabled"
else
    print "feature is disabled"
end
```
