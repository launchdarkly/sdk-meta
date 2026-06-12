---
id: erlang-server-sdk/sdk-docs/features/filedata/flags-from-files
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: File data source configuration example for Erlang.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only

---

```erlang
  ldclient:start_instance("YOUR_SDK_KEY", #{
    file_datasource => true,
    send_events => false,
    file_paths => ["file1.json", "file2.yaml"],
    feature_store => ldclient_storage_map,
    file_auto_update => true,
    file_poll_interval => 1000
  })

  %% In the Erlang SDK automatic reloading uses a polling mechanism.
  %% The default interval is 1000ms, but you can control it with
  %% the file_poll_interval configuration.
```
