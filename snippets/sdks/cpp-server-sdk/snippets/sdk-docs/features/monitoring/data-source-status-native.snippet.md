---
id: cpp-server-sdk/sdk-docs/features/monitoring/data-source-status-native
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Data source status change listener for C++ (server-side) SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```cpp
client.DataSourceStatus().OnDataSourceStatusChange([](server_side::DataSourceStatus status) {
  if (status.State() ==
  server_side::DataSourceStatus::
  DataSourceState::kValid) {
    /* Flag data has been received from LaunchDarkly.*/
  }
});
```
