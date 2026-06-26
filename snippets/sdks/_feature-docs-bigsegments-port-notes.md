# Port notes: /sdk/features/big-segments

Source: `ld-docs-private` `fern/topics/sdk/features/big-segments.mdx`.
10 code blocks extracted into `sdk-docs/features/bigsegments/`
snippets across 7 server-side SDKs (.NET, Go, Java, Node.js x4, PHP,
Python, Ruby). All 10 are bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0).

- **.NET sample** (`dotnet-server-sdk/.../big-segments`): two fixes.
  `BigSegmentsConfigurationBuilder` has no `CacheTime` method — the
  cache-duration option is `ContextCacheTime` (paired with the
  `ContextCacheSize` the sample already used); and the client class in
  the .NET server SDK is `LdClient`, not `LDClient` (the published
  line `new LDClient(config)` is Java-style casing that does not
  compile; the already-ported config/evaluating snippets use
  `LdClient` for the same reason).
- **Node.js SDK v8.x JavaScript sample**
  (`node-server-sdk/.../big-segments-js-v8`): the first import line
  ended with `');` — a stray close-paren left over from converting a
  `require(...)` call to an `import` statement, which is a syntax
  error. Now `import { init } from '@launchdarkly/node-server-sdk';`.

## Validation routing added in this port

- `LaunchDarkly.ServerSdk.Redis` added to the
  `csharp-syntax-only` scaffold's requirements so bodies that
  reference the Redis store integration (`Redis.BigSegmentStore()`)
  compile. The harness `dotnet add package`s each requirements line,
  so the addition is purely additive for existing wrappees.
- `launchdarkly-java-server-sdk-redis-store` (3.1.1, brings jedis
  transitively) added to the jvm harness's synthesized pom so
  `Redis.bigSegmentStore()` resolves for the Java sample.
- `sdkKey` stub field added to the `java-syntax-only` scaffold —
  init-shaped fragments pass an `sdkKey` the docs assume already
  exists (same stub the test-data port adds).

No new scaffolds or validators were needed: the Go, Node, PHP,
Python, and Ruby samples ride the existing parse-only scaffolds,
which don't resolve the store-integration packages
(`ldredis`, `@launchdarkly/node-server-sdk-redis`, Predis, etc.) and
so need no dependency staging.

## Known non-binds

None — every code block on the page is bound to a validator.
