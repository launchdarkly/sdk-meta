---
id: go-server-sdk/sdk-docs/features/flush/flush-interval
sdk: go-server-sdk
kind: reference
lang: go
description: Flush interval configuration example for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
config := ld.Config{
    Events: ldcomponents.SendEvents().FlushInterval(time.Second*10),
}
```
