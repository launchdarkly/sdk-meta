# Port notes: /sdk/features/opentelemetry

Source: `ld-docs-private` `fern/topics/sdk/features/opentelemetry.mdx`
(published as `/sdk/features/opentelemetry-server-side`).
31 code blocks extracted into `sdk-docs/features/otel/` snippets across
7 server-side SDKs plus six shared page-level examples (one
environment-variable export, five OpenTelemetry collector YAML
configurations). All 25 per-SDK snippets are bound to validators; the
six `_shared` examples are documented non-binds (below).

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0).

- **.NET (server) tracing hook with options**
  (`dotnet-server-sdk/.../tracing-hook-options`): the second `using`
  directive read `LaunchDarkly.Sdk.Server.OpenTelemetry` — no such
  namespace exists. The `LaunchDarkly.ServerSdk.Telemetry` package's
  root namespace is `LaunchDarkly.Sdk.Server.Telemetry` (the same
  namespace the basic example on the page already imports). Fixed the
  directive.
- **.NET (server) tracing hook with environment ID**
  (`dotnet-server-sdk/.../tracing-hook-environment-id`): same
  `OpenTelemetry` --> `Telemetry` namespace fix, plus the builder chain
  passed `TracingHook.Builder().EnvironmentId(...)` directly to
  `HookConfigurationBuilder.Add(Hook)` — a `TracingHookBuilder` is not
  a `Hook`, so the sample did not compile. Added the missing
  `.Build()`.
- **Java resource attributes** (`java-server-sdk/.../resource-attributes`):
  the sample called `OpenTelemetrySdk.builder().setResource(resource)`,
  but `OpenTelemetrySdkBuilder` has no `setResource` method — in the
  OpenTelemetry Java SDK a resource is attached to a provider, not to
  the top-level builder. Rewritten to set the resource on
  `SdkTracerProvider.builder()` and register that provider via
  `setTracerProvider(...)`, matching the structure of the Go and
  Python examples on the same page.

## Validation routing added in this port

- `validators/languages/jvm/harness/run.sh` — the synthesized pom now
  also pins `com.launchdarkly:launchdarkly-java-server-sdk-otel:0.2.0`
  and `io.opentelemetry:opentelemetry-sdk:1.40.0`. The tracing-hook
  fragments import `com.launchdarkly.integrations.*`, and the
  `java-syntax-only` scaffold now pre-imports OpenTelemetry SDK types,
  so both artifacts must be on every compile's classpath.
- `java-server-sdk/scaffolds/java-syntax-only` — added
  `import java.util.Collections;` (the tracing-hook fragments call
  `Collections.singletonList` without showing the import) and
  single-type imports for `OpenTelemetry`, `OpenTelemetrySdk`,
  `Resource`, and `SdkTracerProvider` (the resource-attributes
  fragment shows no import lines at all). All previously bound jvm
  wrappees revalidated.
- `dotnet-server-sdk/scaffolds/csharp-syntax-only` — requirements now
  also pin `LaunchDarkly.ServerSdk.Telemetry`, which ships
  `TracingHook` / `TracingHookBuilder` in the
  `LaunchDarkly.Sdk.Server.Telemetry` namespace.
- `validators/languages/shell-install/harness/run.sh` — new
  `composer require <vendor>/<pkg>` dispatch branch. The PHP install
  block invokes bare `composer`, which the image ships only as
  `/opt/composer.phar`; the branch stages a thin `composer` wrapper on
  PATH, initializes a minimal composer project, runs the body, and
  asserts `vendor/<vendor>/<pkg>` exists — the same flow as the
  existing `php composer.phar require` branch.

## Known non-binds

Six `_shared` page-level examples (`sdk-docs/features/otel/...`), all
blocked the same way as the filedata port's shared examples:
`snippets validate` requires `--sdk` and filters on the frontmatter
`sdk:` field, which `_shared` snippets do not carry — no CI row can
currently select them. Per-format blockers on top of that:

- `resource-attributes-collector`, `collector-config-http-multi-env`,
  `collector-config-grpc-multi-env`, `collector-config-http-single-env`,
  `collector-config-grpc-single-env` — no `yaml` validator runtime
  exists yet (a PyYAML-style parse harness would do).
- `resource-attributes-env-var` — a bare `export VAR=...` shell
  fragment; the `shell-install` harness is a package-manager sniff
  followed by an artifact assertion, and a bare export has nothing to
  assert post-run. Same blocker as the web-proxy port's `https-proxy-*`
  fragments.
