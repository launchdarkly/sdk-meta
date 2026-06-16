---
id: cpp-server-sdk/sdk-docs/features/events/track-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Custom event tracking example for C++ (server-side) SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```cpp
/* track the event */
client.Track(context, "example-event-key");

/* track the event and associate data with it */
client.Track(context, "example-event-key", 42);

```
