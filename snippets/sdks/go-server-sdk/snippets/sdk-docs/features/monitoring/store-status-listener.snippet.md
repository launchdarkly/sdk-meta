---
id: go-server-sdk/sdk-docs/features/monitoring/store-status-listener
sdk: go-server-sdk
kind: reference
lang: go
description: Data store status listener channel for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
dataStoreStatusChannel := client.GetDataStoreStatusProvider().AddStatusListener()
go func() {
    for status := range dataStoreStatusChannel {
        fmt.Println("new status is: ", status)
    }
}()
```
