# Port notes: /sdk/features/storing-data (subdirectory)

Source: `ld-docs-private` `fern/topics/sdk/features/storing-data/`
(four pages: `index.mdx`, `redis.mdx`, `dynamodb.mdx`, `consul.mdx`).
56 code blocks extracted into per-page groups
`sdk-docs/features/storing-data/{index,redis,dynamodb,consul}/`
across 11 server-side SDKs. All 56 are bound to validators — no
non-binds (the pages have no iOS/Objective-C sections).

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0).

- **.NET client casing** (`dotnet-server-sdk/.../redis/redis`,
  `dynamodb/dynamodb`, `consul/consul`): `new LDClient(config)` is
  Java-style casing; the .NET server SDK class is `LdClient` (same
  fix the big-segments port applied).
- **.NET Consul name collision** (`dotnet-server-sdk/.../consul/consul`):
  `LaunchDarkly.ServerSdk.Consul` depends on Consul.NET, which
  declares a global `Consul` *namespace*. In any project referencing
  the package, the unqualified `Consul.DataStore()` resolves to that
  namespace, not the `LaunchDarkly.Sdk.Server.Integrations.Consul`
  class (C# checks enclosing-namespace members before using-imports),
  producing CS0234. The call is now fully qualified.
- **C++ v3 native Redis source** (`cpp-server-sdk/.../redis/redis-v3`):
  `LazyLoad().Source(*redis_source)` does not compile —
  `RedisDataSource::Create` yields a `unique_ptr` inside the expected,
  and `Source` takes a `shared_ptr`, which only converts from a
  `unique_ptr` rvalue. Now `Source(std::move(*redis_source))`,
  matching the SDK's own hello-cpp-server-redis example.
- **C++ v3 C-binding Redis source** (`cpp-server-sdk/.../redis/redis-v3-c`):
  two fixes. The final build step used the *client*-side types
  (`LDClientConfig config; LDClientConfigBuilder_Build(...)`) on the
  server config builder; corrected to `LDServerConfig` /
  `LDServerConfigBuilder_Build` per the server C binding header. And
  `LDServerLazyLoadBuilder_SourcePtr(lazy_builder, result.source)`
  passes an `LDServerLazyLoadRedisSource` where the builder takes an
  `LDServerLazyLoadSourcePtr` — an incompatible-pointer constraint
  violation; the SDK's own hello-c-server-redis example inserts the
  `(LDServerLazyLoadSourcePtr)` cast, now applied here too.
- **Haskell Redis store** (`haskell-server-sdk/.../redis/redis`):
  `configSetStoreBackend` takes `Maybe PersistentDataStore` in the
  v4 SDK, but `makeRedisStore` produces a bare `PersistentDataStore`;
  now `configSetStoreBackend (Just backend)`. (The index page's
  generic sample is unchanged: its placeholder constructor can be
  inferred to return the `Maybe`.)
- **Lua v1.x client init** (`lua-server-sdk/.../index/persistent-store-v1x`,
  `redis/redis-v1x`): the `clientInit` call closed with `)}` instead
  of `})` — a syntax error luac rejects.
- **Node.js v8 Redis require** (`node-server-sdk/.../redis/redis-js-v8`):
  `@launchdarkly/node-server-sdk-redis` has no callable module
  export — `RedisFeatureStore` is a named export — so
  `const RedisFeatureStore = require(...)` followed by
  `RedisFeatureStore({...})` throws at runtime. Now destructured,
  matching the package's index.ts and the page's own TypeScript and
  DynamoDB variants.
- **PHP phpredis integration** (`php-server-sdk/.../redis/redis-phpredis`):
  the page directs readers to v2 of `server-sdk-redis-phpredis`, whose
  `featureRequester` requires a configured `Redis` client instance as
  the first argument (the array-options form is the removed v1.x
  shape). Rewritten to construct and connect a `Redis` client first,
  per the integration's README.
- **Ruby client class** (`ruby-server-sdk/.../index/persistent-store`,
  `redis/redis`, `dynamodb/dynamodb`, `consul/consul`):
  `LaunchDarkly::Client` does not exist in the Ruby SDK; the client
  class is `LaunchDarkly::LDClient` (as the index page's own
  DataSystem sample uses).

## Validation routing added in this port

- `cpp-server-sdk/scaffolds/cpp-syntax-only-redis` — variant of
  `cpp-syntax-only` that pre-includes the Redis source headers
  (native + C binding) at file scope and sets `CPP_REDIS=1` via
  `validation.env`. The cpp-server harness now configures cpp-sdks
  with `-DLD_BUILD_REDIS_SUPPORT=ON` and links
  `launchdarkly::server_redis_source` when that flag is set; the
  Dockerfile prewarm builds the redis source target so per-validate
  cycles stay ccache-warm.
- `erlang-server-sdk/scaffolds/erlang-syntax-only-stmts` — variant of
  `erlang-syntax-only` for statement-sequence fragments that carry
  their own clause-terminating dot (the plain scaffold appends one,
  which would double-terminate these bodies).
- `haskell-server-sdk/scaffolds/haskell-syntax-only-module` —
  compile-only scaffold for module-shaped fragments that define their
  own `main = do`. The haskell-server harness gained a
  `SNIPPET_CHECK=parse` path (same convention as erlang-server) that
  stops after a clean `cabal build`, since running the body's `main`
  would try to reach a real database. The scaffold completes the
  fragments' trailing `client <- makeClient config` bind (a doc idiom)
  with an indented `pure ()` so the do-block ends in an expression
  without altering the published body, and supplies an untyped
  `makeYourBackendInterface = undefined` stub.
- Store-integration dependencies staged for the compile-based
  validators:
  - jvm harness pom: `launchdarkly-java-server-sdk-redis-store` 3.1.1
    (same addition as the big-segments port), plus
    `launchdarkly-java-server-sdk-dynamodb-store` 5.0.0 and
    `launchdarkly-java-server-sdk-consul-store` 5.0.0.
  - `csharp-syntax-only` requirements: `LaunchDarkly.ServerSdk.Redis`
    (same addition as the big-segments port), plus
    `LaunchDarkly.ServerSdk.DynamoDB` and
    `LaunchDarkly.ServerSdk.Consul`.
  - haskell-server Dockerfile cabal deps: `hedis` and
    `launchdarkly-server-sdk-redis-hedis`.
- `cpp-server-v2-c` validator: new stub `<launchdarkly/store/redis.h>`
  (LDRedisConfigNew / LDStoreInterfaceRedisNew) and
  `LDConfigSetFeatureStoreBackend` + `struct LDStoreInterface` added
  to the stub api.h; `cpp-syntax-only-v2-c` pre-includes the store
  header so the fragment's own in-body `#include` is an include-guard
  no-op.
- Stub-surface extensions: `sdkKey` + nested `SomeDatabaseName`
  placeholder class on `csharp-syntax-only`; `sdkKey`, `storeOptions`
  + nested `SomeDatabaseName` placeholder class and
  `com.launchdarkly.sdk.server.integrations.*` / `java.net.URI` /
  `java.net.URL` pre-imports on `java-syntax-only` (the index page's
  DataSystem sample references `Redis` with only the base import, and
  the redis/consul samples use `URI.create` / `new URL` without their
  own imports); `sdk_key` and a file-scope `YourDatabaseIntegration()`
  placeholder returning `LazyLoadBuilder`'s source pointer type on
  `cpp-syntax-only`.

## Known non-binds

None — every code block across the four pages is bound to a validator.
