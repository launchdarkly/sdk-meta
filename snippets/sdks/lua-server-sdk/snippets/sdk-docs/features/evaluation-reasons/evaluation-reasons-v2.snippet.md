---
id: lua-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Flag evaluation reason example for Lua SDK v2.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local details = client:boolVariationDetail(context, "example-flag-key", false);

-- inspect details here
if details.reason.kind == "ERROR" and details.reason.errorKind == "FLAG_NOT_FOUND" then
end
```
