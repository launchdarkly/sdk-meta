# Port notes: /sdk/features/logging

Source: `ld-docs-private` `fern/topics/sdk/features/logging.mdx`.
49 code blocks extracted into `sdk-docs/features/logging/` snippets
across 20 SDKs. All but one (iOS Objective-C) are bound to
validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0 for the indented CodeBlock containers).

- **Go logging example** (`go-server-sdk/.../logging-v6`): the import
  line read `ldlog "github.com/launchdarkly/go-sdk-common/v3"`, but
  the `ldlog` package lives in the `ldlog` subdirectory of that
  module; corrected to
  `ldlog "github.com/launchdarkly/go-sdk-common/v3/ldlog"`.
- **Node.js (client-side) TypeScript example**
  (`node-client-sdk/.../logging-ts`): the body imports the package as
  `LaunchDarkly` (`import * as LaunchDarkly from ...`) but then
  called `ld.basicLogger(...)` on an undefined `ld` binding;
  corrected to `LaunchDarkly.basicLogger(...)`.
- **Roku SceneGraph custom logger**
  (`roku-client-sdk/.../custom-logger-scenegraph-brs`): the observer
  loop tested `field = "log"` against an undefined `field` variable
  (always `Invalid`, so the handler never ran). Changed to
  `msg.getField()`, the changed-field name from the `roSGNodeEvent`.

## Validation routing added in this port

- `cpp-server-sdk/scaffolds/cpp-syntax-only-toplevel` and
  `cpp-client-sdk/scaffolds/cpp-client-syntax-only-toplevel` --
  file-scope splice variants for the custom-log-backend fragments,
  which are themselves top-level declarations (an `ILogBackend`
  subclass; `static` C-binding callback definitions) and cannot live
  inside the nested `_wrappee()` body.
- `cpp-server-sdk/scaffolds/cpp-syntax-only-v2-c-toplevel` --
  file-scope splice variant for the v2 C server SDK custom-logger
  function definition (C forbids `static` function definitions inside
  another function).
- `roku-client-sdk/scaffolds/roku-syntax-only-toplevel` -- file-scope
  splice variant for the SceneGraph `CustomLogger.brs` fragment (two
  top-level function definitions).
- `roku-client-sdk/scaffolds/roku-syntax-only-exprlist` -- the
  supported-log-levels fragment is a list of bare
  `LaunchDarklyLogLevels().x` expressions; BrightScript forbids a
  property access as an expression statement, so this variant splices
  the body as array-literal elements (a legal expression context).
- `roku-client-sdk/scaffolds/roku-xml-syntax-only` + the new `xml`
  validator (`validators/languages/xml/`) -- `xmllint --noout`
  well-formedness pass for the SceneGraph `CustomLogger.xml`
  component fragment. Parse-only; no SceneGraph schema.
- Stub-surface extensions:
  - cpp v3 nested scaffolds (server + client): `<memory>` and the
    logging interface headers at file scope; a `CustomLogger`
    backend stub plus `enabled`/`write` callback stubs so the
    install-a-custom-logger fragments (which reference symbols
    defined in a preceding fragment on the page) compile standalone.
    The server scaffold additionally lifts
    `launchdarkly::server_side::config::builders` inside `_wrappee()`
    because the install fragment uses unqualified `LoggingBuilder`,
    relying on the alias the page's basic-logging fragment declared.
  - `cpp-syntax-only-v2-c`: `myCustomLogger` stub for the v2
    install fragment.
  - v2 stub header (`validators/languages/cpp-server-v2-c/api.h`):
    `LDLogLevel` enum, `LDBasicLogger`, `LDConfigureGlobalLogger`,
    `LDLogLevelToString`, mirroring c-sdk-common's
    `<launchdarkly/logging.h>`.
  - android `java-syntax-only` + `kotlin-syntax-only`:
    `timber.log.Timber` import (the Timber fragments call
    `Timber.plant(...)` without showing the import; the validator
    project already ships the Timber dependency).
  - `csharp-syntax-only`: `loggerFactory` stub (`dynamic`) for the
    Microsoft.Extensions.Logging adapter fragment.

## Known non-binds

- `ios-client-sdk/.../logging-objc` -- no Objective-C parse scaffold
  exists; the iOS validator is the macOS-only native harness (same
  blocker as the evaluating and evaluation-reasons ports' objc
  snippets). Wiring it up requires either an Objective-C target in
  the xcodegen scaffold or a clang -fsyntax-only stub harness.
