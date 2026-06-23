# Port notes: /sdk/features/monitoring

Source: `ld-docs-private` `fern/topics/sdk/features/monitoring.mdx`.
35 code blocks extracted into `sdk-docs/features/monitoring/`
snippets across 15 SDKs. All 35 are bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0).

- **Android status listener (Java)**
  (`android-client-sdk/.../status-listener-java`): `LDClient.get()`
  throws the checked `LaunchDarklyException`, and the sample calls it
  inside an override of `Activity.onCreate`, which cannot declare a
  `throws` clause — the sample cannot compile in any context. Wrapped
  the `get()` call in try/catch.
- **Android status listener (Kotlin)**
  (`android-client-sdk/.../status-listener-kotlin`): the sample
  declares `client` as nullable (`LDClient?`) but registers the
  listener with a plain `client.registerStatusListener(...)` call —
  Kotlin requires a safe or non-null-asserted call on a nullable
  receiver. Changed to `client?.registerStatusListener(...)`,
  matching the sample's own `client?.unregisterStatusListener(...)`
  in `onDestroy`.
- **C++ (client-side) C binding listener**
  (`cpp-client-sdk/.../data-source-status-c-create-listener`):
  `LDDataSourceStatusListener_Init(listener)` passed the struct by
  value; the real binding takes
  `struct LDDataSourceStatusListener *` — changed to `&listener`
  (the server-side block on the same page already uses `&listener`).
- **C++ (server-side) native listener**
  (`cpp-server-sdk/.../data-source-status-native`): the sample
  referenced `server_side::data_sources::DataSourceStatus`; the
  server SDK declares `DataSourceStatus` directly in
  `launchdarkly::server_side` — there is no `data_sources`
  sub-namespace on the server side (unlike the client SDK).
- **Flutter v3.x connection information**
  (`flutter-client-sdk/.../connection-information-v3`):
  `LDClient.getConnectionInformation()` returns
  `LDConnectionInformation?` (nullable) in v3.x; the sample assigned
  it to a non-nullable local and read fields off it. Declared the
  local nullable and made the field reads null-aware.

## Validation routing added in this port

- `cpp-client-sdk/scaffolds/cpp-client-syntax-only-toplevel` and
  `cpp-server-sdk/scaffolds/cpp-syntax-only-toplevel` — file-scope
  splice variants for the C-binding "define a callback" fragments.
  C++ forbids function definitions inside another function, so these
  cannot ride the default `_wrappee()` splice.
- `dotnet-client-sdk/scaffolds/csharp-client-syntax-only-typed` and
  `dotnet-server-sdk/scaffolds/csharp-syntax-only-typed` — variants
  whose `client` stub is the real `LdClient` instead of `dynamic`.
  C# rejects a lambda as an operand of a dynamically dispatched
  operation (CS1977), so the `StatusChanged += (sender, status) => …`
  event-subscription fragments cannot compile against the default
  scaffolds' `dynamic` stub. Typing the stub also makes the event
  subscription type-check against the real SDK surface.
- Stub-surface extensions:
  - `_AnyStatusProvider` (variadic `OnDataSourceStatusChange`)
    returned by a new `DataSourceStatus()` member on `_AnyClient` in
    both cpp syntax-only scaffolds, for the native listener
    fragments.
  - File-scope `OnDataSourceStatusChanged` callback stubs in both
    cpp syntax-only scaffolds — the create-listener fragments assign
    a callback that the docs define in a separate code block.
  - `sdk` (client handle) and `connection` (`LDListenerConnection`)
    ambient locals in the cpp syntax-only scaffolds — the C-binding
    fragments reference both as pre-existing.
  - `LDStatus` enum + `LDSetClientStatusCallback` + `<stdio.h>` on
    the `cpp-client-v2-c` validator's stub `<launchdarkly/api.h>`
    for the C SDK v2.x status-callback fragment. The fragment's
    top-level callback definition plus registration statement
    compiles through the existing `_wrappee()` splice because gcc
    accepts nested function definitions without `-pedantic`.
  - android `java-syntax-only`: Timber import, and the spliced body
    now sits in a try/catch so statement-shaped fragments calling
    checked-throwing APIs (`LDClient.get()`) have a legal home
    (mirrors the java-server scaffold).
  - android `kotlin-syntax-only`: `android.app.Activity` import (the
    status-listener fragment declares a local
    `class MainActivity : Activity()`), Timber import, and
    `LDFailure.FailureType.*` enum-entry import (Kotlin does not
    resolve unqualified enum entries in `when` branches from the
    subject's type).
  - java-server `java-syntax-only`:
    `com.launchdarkly.sdk.server.interfaces.*` import so
    `DataSourceStatusProvider.Status` resolves unqualified.

## Known non-binds

None — every block on the page is bound. (This page has no
Objective-C block; the iOS Swift snippet rides the existing
`swift-syntax-only` scaffold, which only the CI macOS row can
execute.)
