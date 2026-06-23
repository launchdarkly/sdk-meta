---
id: cpp-client-sdk/sdk-docs/features/monitoring/data-source-status-native
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Data source status change listener for C++ (client-side) SDK v3.0 (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
client.DataSourceStatus().OnDataSourceStatusChange([](client_side::data_sources::DataSourceStatus status) {
  if (status.State() ==
  client_side::data_sources::DataSourceStatus::
  DataSourceState::kValid) {
    /* Flag data has been received from LaunchDarkly.*/
  }
});
```
