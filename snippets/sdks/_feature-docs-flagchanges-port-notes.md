# Port notes: /sdk/features/flag-changes

Source: `ld-docs-private` `fern/topics/sdk/features/flag-changes.mdx`.
39 code blocks extracted into `sdk-docs/features/flagchanges/`
snippets across 16 SDKs. All but one (iOS Objective-C) are bound to
validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0 for the C-binding blocks that sit inside a numbered list).

- **Android Java listener registration**
  (`android-client-sdk/.../flag-changes-java`): the
  `onFeatureFlagChange` override calls `LDClient.get()`, which throws
  the checked `LaunchDarklyException`; an interface-override cannot
  add a throws clause, so the body cannot compile in any caller
  context. Wrapped the call in try/catch. The outer
  `LDClient.get().registerFeatureFlagListener(...)` line is left
  verbatim — it compiles in a caller method that may throw, which the
  scaffold now models (see routing below).
- **C++ (client-side) C binding, create listener connection**
  (`cpp-client-sdk/.../flag-changes-cpp-c-v3-0-create-connection`):
  the published block called `LDFlagListener_Init(listener)` by
  value; the real header takes the struct by pointer
  (`LDFlagListener_Init(struct LDFlagListener*)`). Changed to
  `LDFlagListener_Init(&listener);`.
- **.NET (server-side) flag-change subscription**
  (`dotnet-server-sdk/.../flag-changes`): the lambda named its
  parameter `event`, a reserved keyword in C# — a syntax error.
  Renamed to `eventArgs`, matching the SDK's own `IFlagTracker` doc
  example.
- **.NET (server-side) v7.0 flag-value-change subscription**
  (`dotnet-server-sdk/.../flag-value-changes-v7`): same reserved
  `event` parameter name; same `eventArgs` rename.

Left as-is on purpose: the JavaScript v4.0 block's `defaultValue` is
a deliberate fill-this-in placeholder (same as `YOUR_SDK_KEY`-style
literals), not an undefined-variable bug, and the JS validator's
parse-only pass accepts it.

## Validation routing added in this port

- `cpp-client-sdk/scaffolds/cpp-client-syntax-only-toplevel` —
  file-scope sibling of `cpp-client-syntax-only` for fragments that
  are themselves top-level declarations (the C-binding callback
  definition); C++ has no local free functions, so the nested-block
  splice cannot host them. Same `cpp-client` validator.
- `dotnet-client-sdk/scaffolds/csharp-client-syntax-only-typed` —
  real-typed `LdClient` stub sibling of `csharp-client-syntax-only`.
  C# forbids lambdas as operands of dynamically dispatched
  operations, so the event-pattern body
  (`client.FlagTracker.FlagValueChanged += (sender, eventArgs) => ...`)
  cannot compile against the dynamic stub. Same `dotnet-client`
  validator container.
- `ruby-server-sdk/scaffolds/ruby-syntax-only-block` — splices the
  body inside a never-called proc instead of `def _wrappee`; Ruby
  rejects `class` definitions inside a method body at parse time,
  and both Ruby blocks on this page define a `Listener` class.
- `android-client-sdk/scaffolds/java-syntax-only` restructure: the
  body now splices into a never-invoked `_wrappee() throws Exception`
  instance method instead of the `onCreate` override, because
  lifecycle overrides cannot declare the checked
  `LaunchDarklyException` that `LDClient.get()` throws. All
  previously bound wrappees re-validated against the new shape.
- Stub-surface extensions (all additive):
  `FlagNotifier()`/`_AnyNotifier`/`_AnyConnection`, an
  `OnFlagChange` callback stub, and `sdk`/`connection` locals on
  `cpp-client-syntax-only`; `listener` field and `java.util.List`
  import on the android `java-syntax-only` scaffold; top-level
  `listener` stub on `kotlin-syntax-only`; `sub` stub on
  `flutter-syntax-only`; `listener` stub on
  `flutter-syntax-only-v3`.

No CI workflow changes: every new binding routes through a runtime
that already has a matrix row, and no snippet targets the v4-era
android jvm split.

## Known non-binds

- `ios-client-sdk/.../flag-changes-objc` — no Objective-C parse
  scaffold exists; the iOS validator is the macOS-only native harness
  (same blocker as the evaluating port's objc snippet). Wiring it up
  requires either an Objective-C target in the xcodegen scaffold or a
  clang -fsyntax-only stub harness.
