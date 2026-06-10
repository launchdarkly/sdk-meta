---
id: lua-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-v1x
sdk: lua-server-sdk
kind: reference
lang: lua
description: Flag evaluation reason example for Lua SDK v1.x.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local details = client:boolVariationDetail(user, "example-flag-key", false);

-- inspect details here
if details.reason.kind == "ERROR" and details.reason.errorKind == "FLAG_NOT_FOUND" then
end
```
