---
id: cpp-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-alias-events-3-0-syntax-working-with-associated-contexts
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "3.0 syntax, working with associated contexts in section \"Understanding changes to alias events\""
---

```cpp
/* before end user logs in */
auto context = ContextBuilder()
  .Kind("device", "example-device-key")
  .Build();


/* after end user logs in */
auto updated_context = ContextBuilder(context)
  .Kind("user", "example-user-key")
  .Name("Sandy")
  .Build();

client.IdentifyAsync(updated_context);
```
