---
id: go-server-sdk/sdk-docs/initialize-the-client-go-sdk-using-ldclient-and-default-configuration
sdk: go-server-sdk
kind: reference
lang: go
description: "Go SDK, using LDClient and default configuration in section \"Initialize the client\""
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
client, _ := ld.MakeClient("YOUR_SDK_KEY", 5*time.Second)
```
