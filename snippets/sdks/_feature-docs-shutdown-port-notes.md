# Port notes: /sdk/features/shutdown

Source: `ld-docs-private` `fern/topics/sdk/features/shutdown.mdx`.
27 code blocks extracted into `sdk-docs/features/shutdown/` snippets
across 22 SDKs. All but two (iOS Objective-C, Erlang) are bound to
validators.

## Content corrections

None. Every body is verbatim from the MDX. The shutdown calls are all
one-liners against APIs that still exist as published (`Dispose()`,
`close()`/`close`, `LDClientSDK_Free` / `LDServerSDK_Free` /
`LDClientClose`, `ldclient.get().close()`, `close client`,
`ldclient:stop_*`), verified against the local SDK checkouts.

## Validation routing added in this port

Every bound snippet routes through a scaffold that already existed
before this port:

- `*-syntax-only` scaffolds for the single-call fragments (the bodies
  reference a bare `client`, which every scaffold already stubs).
- `cpp-client-syntax-only-v2-c` / `cpp-client-syntax-only-v2-cpp` /
  `cpp-syntax-only-v2-c` for the C SDK v2.x blocks; the stub
  `<launchdarkly/api.h>` / `<launchdarkly/api.hpp>` headers already
  declare `LDClientClose` and `LDClientCPP::close`.
- `android-client-sdk/scaffolds/java-syntax-only` (v5 aar) for the
  Android Java block — `client.close()` is valid on the v5 API, so the
  v4 jvm-routed stub scaffold is not needed here.

One scaffold extension (additive):

- `android-client-sdk/scaffolds/java-syntax-only` — the unreachable
  `if (false)` body is now additionally wrapped in
  `try { … } catch (Exception e) { }` (mirroring the
  `csharp-syntax-only` scaffold's shape). The Android `client.close()`
  fragment calls a checked-exception API (`java.io.Closeable.close()`
  throws `IOException`), and the body's host method `onCreate`
  overrides `Activity.onCreate`, so a `throws` clause can't be added
  there. Catching `Exception` is legal even when the body throws no
  checked exception, so previously bound fragments are unaffected
  (re-validated `evaluation-reasons/print-reason-java` to confirm).

## Known non-binds

- `ios-client-sdk/.../shutdown-objc` — no Objective-C parse scaffold
  exists; the iOS validator is the macOS-only native harness (same
  blocker as the evaluating and evaluation-reasons ports' objc
  snippets). Wiring it up requires either an Objective-C target in the
  xcodegen scaffold or a clang -fsyntax-only stub harness.
- `erlang-server-sdk/.../shutdown` — the block shows three
  *alternative* calls (`stop_all_instances()`, `stop_instance()`,
  `stop_instance(my_instance)`) separated by comment lines, which is
  not a valid Erlang expression sequence inside the parse scaffold's
  single function body (Erlang requires `,` separators, and adding
  them would misrepresent the alternatives as one sequence). Same
  shape, and same treatment, as the unbound
  `erlang-server-sdk/sdk-docs/features/config/index` snippet. Wiring
  it up requires either splitting the doc block into one block per
  alternative (a docs content change) or teaching the erlang-server
  harness a pre-rewrite that stages each blank-line-separated
  expression as its own function clause.
