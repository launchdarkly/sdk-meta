---
id: go-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode
sdk: go-server-sdk
kind: reference
lang: go
description: Daemon mode configuration example for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
config := ld.Config{
    DataStore:  ldcomponents.PersistentDataStore(
      examplepackage.DataStore().SomeStoreOptions(),
    ),
    DataSource: ldcomponents.ExternalUpdatesOnly(),
}

```
