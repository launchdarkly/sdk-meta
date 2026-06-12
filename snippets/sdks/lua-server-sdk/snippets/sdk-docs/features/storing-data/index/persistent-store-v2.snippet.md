---
id: lua-server-sdk/sdk-docs/features/storing-data/index/persistent-store-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Lazy-load data source configuration example for Lua SDK v2.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local config = {
    dataSystem = {
        lazyLoad = {
            source = makeYourSource()
        }
    }
}
```
