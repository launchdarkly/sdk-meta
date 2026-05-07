---
id: go-server-sdk/ai-configs/context
sdk: go-server-sdk
kind: context
lang: go
file: go-server-sdk/ai-configs/context.txt
description: Build an evaluation context for go-server-sdk AI Configs.
---

```go
context := ldcontext.NewBuilder("context-key-123abc").
    Kind("user").
    Name("Sandy Smith").
    SetString("email", "sandy@example.com").
    SetValue("groups", ldvalue.ArrayOf(
      ldvalue.String("Google"), ldvalue.String("Microsoft"))).
    Build()
```
