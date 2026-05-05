# sdk-docs port notes

Snippets ported from ld-docs MDX into `snippets/sdks/<sdk>/snippets/sdk-docs/`
that the validator harness can't currently exercise are documented here.
Each entry describes the structural reason the snippet has no
`validation:` block and what would need to change for it to become
bindable.

This file is the sdk-docs analogue of `_sdk-info-port-notes.md`. The
sdk-info port had a small handful of structurally unbindable cases; the
sdk-docs slice has many more because the source MDX often shows
fragments that aren't standalone-executable (env-var exports for shell,
deprecated-version API surfaces no longer in the registry, etc.).

## Shell snippets that aren't a package-manager install

**Severity**: low (informational fragments)

**Snippets affected**: every per-SDK `proxy-env-mac`, `proxy-env-windows`,
and similar shell fragments under `sdk-docs/` that just `export VAR=…`
or `set VAR=…` rather than running a package manager.

**Why unbindable**: the `shell-install` validator is a package-manager
sniff (npm/pnpm/yarn/pip/go/bower) followed by an artifact assertion
(node_modules, go.mod, etc.). A bare `export HTTPS_PROXY=…` doesn't fit
that contract — there's nothing to assert post-run, and routing to a
generic "run this in /bin/sh and check exit 0" runner would silently
green-light malformed bodies. These snippets are byte-equality-verified
through marker hashes already, which is the strongest check available
for a fragment that doesn't have a runtime success-line.

## Multi-line `pip install` snippets

**Severity**: low

**Snippets affected**: `python-server-sdk/sdk-docs/install` (installs
both the SDK and the optional observability plugin in one block).

**Why unbindable**: the `shell-install` harness's `last_pkg` extractor
runs awk per-line and prints the last token of every non-empty line, so
a body with two `pip install` commands produces a multi-line `$pkg`
that breaks the post-install `pip show $pkg` assertion. The single-line
`sdk-info/install-pip.snippet.md` covers the SDK install, so the docs
fragment doesn't add validatable surface that isn't already covered.

## .NET client SDK v3.x fragments

**Severity**: low (older API surface)

**Snippets affected**:
`dotnet-client-sdk/sdk-docs/initialize-the-client-net-sdk-v3-0-c`,
`dotnet-client-sdk/sdk-docs/initialize-the-client-net-sdk-v3-x-c`,
`dotnet-client-sdk/sdk-docs/using-the-relay-proxy-c`.

**Why unbindable**: these reference v3-shape `LdClient.Init(string,
Context, TimeSpan)` and the old `Configuration.Builder(...)
.ServiceEndpoints(...)` chain that changed in v4 (the
`AutoEnvAttributes` parameter is now mandatory). The scaffold pulls
the latest LaunchDarkly.ClientSdk, so these v3 overloads no longer
exist on the type. Pinning v3 in a parallel csproj would let them
build, but would diverge from the canonical "what should I install"
direction the rest of the bindings reflect.

## C++ SDK v2.x fragments (cpp-client, cpp-server)

**Severity**: medium (fragments still byte-checked, but no compile)

**Snippets affected**: every `*-c-sdk-v2-x-*` snippet under
`cpp-client-sdk/sdk-docs/` and `cpp-server-sdk/sdk-docs/`. v2.x used a
C-style API (`LDClientCPP`, `LDClientInit`, `LDConfigNew`,
`LDUserNew`, `<launchdarkly/api.h>`) that the cpp-sdks v3 build does
not ship — even the headers were renamed under
`launchdarkly/client_side/`. The validator Dockerfile pins
`cpp-sdks v3.x` (matching the current docs guidance), so embedding a
v2.x identifier in any function body fails compilation regardless of
whether the function is invoked.

**Why unbindable**: the cpp-sdks repo no longer ships a v2.x branch in
the Dockerfile-pinned form, and re-pinning v2 in addition would either
require a parallel Docker image or a meaningful refactor of the
existing image. The fragments remain byte-equality-verified through
the marker hash machinery; the v3 fragments under the same folder are
fully bound to the cpp-syntax-only scaffold.

## Apex SDK fragments

**Severity**: low (no validator infrastructure)

**Snippets affected**: every snippet under `apex-server-sdk/sdk-docs/`.

**Why unbindable**: there is no Apex/Salesforce validator container —
PR #414's port notes (sdk-info) already documented this gap for the
install/init path. Apex requires the Salesforce CLI and a scratch org
to compile, neither of which fits the existing Docker-based validator
mold. The same applies to the sdk-docs slice; there is no parse-only
fallback because no container ships an Apex parser.

## Roku SDK fragments

**Severity**: low (no validator infrastructure)

**Snippets affected**: every snippet under `roku-client-sdk/sdk-docs/`.

**Why unbindable**: there is no Roku BrightScript validator. The roku
syntax-only scaffold exists as a stub (PR #418) but no Docker image
ships a BrightScript parser. Same disposition as Apex above; the
fragments are byte-checked via marker hashes.

## ASP.NET / Global.asax C# fragments (.NET server SDK)

**Severity**: low (init pattern; not the canonical hello-world flow)

**Snippets affected**:
`dotnet-server-sdk/sdk-docs/initialize-the-client-initialize-net-sdk-v8-10-with-asp-net-core-p1`,
`…-p2`.

**Why unbindable**: both fragments paste into specific .NET hosting
contexts (top-level `Program.cs` with implicit `args` for ASP.NET
Core, or `Global.asax.cs`'s class body for ASP.NET Framework) where
the surrounding scaffold provides the context the body relies on.
The `csharp-syntax-only` scaffold wraps the wrappee in a method body,
so `args` is out of scope (it's a parameter of the implicit
top-level `Main`) and `protected` method definitions can't nest
inside another method. A correct binding would need either two new
ASP.NET Core / ASP.NET Framework scaffolds that supply the right
shape, or a runtime that compiles the body into the correct
hosting context.

## "Pick one of these" multi-style import fragments

**Severity**: low (each individual import style is correct in isolation)

**Snippets affected**:
`js-client-sdk/sdk-docs/install-the-sdk-javascript-and-typescript-js-sdk-v3-7`,
`js-client-sdk/sdk-docs/install-the-sdk-javascript-and-typescript-js-sdk-v4-x`.

**Why unbindable**: the body shows three parallel ways to import
`LDClient` (CommonJS `require`, ES module `import`, and a TS import)
in the same code block. They're meant to be alternatives, not
co-existing imports — but the validator sees one file with three
declarations of the same identifier and rejects the redeclaration.
Splitting the snippet into three separately-bindable variants would
diverge from the source-MDX fingerprint; the right fix lives in
ld-docs (`{/* prettier-ignore */}` comments aren't enough; the docs
should show one canonical example).

## HTML `<script src="…">` fragments

**Severity**: low (declarative; not executable JS)

**Snippets affected**: every `*-html` snippet under
`js-client-sdk/sdk-docs/`, plus the `make-the-sdk-available-with-a-script-tag-…`
pair.

**Why unbindable**: these are HTML script tags — declarative loaders
for an external bundle, not JavaScript. The js-syntax-only scaffold
writes the wrappee body to a `.ts` file that tsdown attempts to
parse; raw HTML isn't valid TS. Each fragment is byte-equality-
verified through the marker hash machinery, which is the strongest
guarantee available for a static asset reference.

## Erlang sdk-docs fragments

**Severity**: medium

**Snippets affected**: every snippet under `erlang-server-sdk/sdk-docs/`.

**Why unbindable**: the `erlang-server` Docker validator's harness
(`run.sh`) hard-codes a synthesized rebar3 `eval` expression that
calls `hello_erlang_server:get(<<flag>>, false, <<key>>)` — it
expects the staged module to define a gen_server that exports
`get/3`. The syntax-only scaffold synthesizes `snippet` with `main/0`
instead, so the harness's eval expression fails to find the expected
module. Every erlang sdk-docs fragment also has shape issues unique
to Erlang grammar (multi-statement bodies need `,` between
expressions; trailing `.` is mandatory at module top-level; `import`
isn't a thing — Erlang uses module qualifiers). To wire these up
cleanly, the harness needs either (a) a parse-only mode keyed on the
scaffold (call `compile:file/2` instead of running eval) or (b) a
new gen_server-shaped scaffold whose body contributes the `get/3`
function clauses. Both are larger refactors than this PR's scope.
Two of the seven sdk-docs fragments are also mistagged in the source
MDX (`get-started-erlang` and `get-started-erlang-2` are
rebar.config / .app.src declarative blocks, and `get-started-elixir`
is Elixir mistagged as Erlang); even with the harness fix above,
those three would need separate routing.

## Android client SDK fragments (java + kotlin)

**Severity**: medium

**Snippets affected**: every snippet under `android-client-sdk/sdk-docs/`.

**Why unbindable**: the existing `android-client-sdk/scaffolds/android-syntax-only`
scaffold routes through the `jvm` validator, which fetches
`launchdarkly-java-server-sdk` from Maven Central — that's the
*server* SDK, not the *android client* SDK. The android client SDK
(`com.launchdarkly:launchdarkly-android-client-sdk`) is published as
an `aar` to Google's Maven, not a plain jar to Maven Central, so
`mvn` can't resolve it without a different repository configuration
plus the Android Gradle Plugin shim. The android sdk-docs fragments
also include kotlin sources, which the jvm/maven validator doesn't
compile. To wire these up cleanly, a parallel
`android-client-validator` Docker image with the Android SDK and
Gradle (the same shape as the existing `android-client-sdk`
sdk-info init validator) would be needed, plus a kotlin-aware
syntax-only scaffold. That's larger than this PR's scope.

The two `install-the-sdk-gradle-*` snippets are Gradle
`implementation '…'` declarations and would be Bucket C even with
the validator above (declarative dependency strings, see XML/Maven
section).

## XML / Maven / Gradle declaration fragments

**Severity**: low (declarative; not executable)

**Snippets affected**: `java-server-sdk/sdk-docs/install-the-sdk-xml`,
`java-server-sdk/sdk-docs/using-the-java-sdk-in-osgi-xml`,
`android-client-sdk/sdk-docs/install-the-sdk-android-sdk-v3-x-android-3-7-x-and-later-android-sdk-3-7-or-later-gradle`
and similar declarative `<dependency>` / `implementation '…'` snippets.

**Why unbindable**: there is no XML parse-and-resolve validator that
asserts a Maven coordinate or Gradle dependency string is reachable;
the closest analogue is the `shell-install` harness, which only works
on package-manager invocations (npm/pnpm/yarn/pip/go/bower). A
declarative `<dependency>` block doesn't fit that contract — the
package is fetched by Maven, not by a CLI invocation we can sniff. The
content is byte-equality-checked through the marker hash machinery,
which is the strongest guarantee available for a declarative snippet.

## iOS Objective-C fragments

**Severity**: medium

**Snippets affected**: every `*-objective-c*` file under
`ios-client-sdk/sdk-docs/`, plus `import-the-sdk-objective-c`.

**Why unbindable**: the swift-syntax-only scaffold compiles Swift,
not Objective-C. Wiring up Objective-C would need a separate
swift-syntax-only-objc scaffold (or harness branch on the
fragment's `lang:`) plus a project.yml that targets an Objective-C
or mixed Swift/Objc app. None of those exist today; the underlying
ios-client SDK supports both languages but the validator harness
only stages Swift sources.

## iOS package-manifest fragments

**Severity**: low (declarative; not executable)

**Snippets affected**:
`ios-client-sdk/sdk-docs/use-carthage-cartfile`,
`ios-client-sdk/sdk-docs/use-cocoapods-podfile`,
`ios-client-sdk/sdk-docs/use-the-swift-package-manager-package-swift`.

**Why unbindable**: Cartfile, Podfile, and Package.swift fragments
are package-manager declarations — same disposition as the XML /
Maven dependency snippets above. The ios-install validator
(separate from ios-client) handles the install flow but only for
the canonical sdk-info install snippets, not arbitrary doc
fragments. These are byte-equality-checked through the marker
hash.

## iOS / Android UIKit / Activity fragments

**Severity**: medium (fragments compile-checked but not run)

**Snippets affected**: ios-client and android-client `sdk-docs/`
snippets that reference `UIViewController`, `Activity`, `application:`
lifecycle hooks, etc.

**Why unbindable** *(sometimes)*: the syntax-only scaffolds compile
the wrappee inside a stub function so unbound symbols don't trip the
parser; symbols like `self.view` or `getApplicationContext()` that
read like instance-method calls compile cleanly inside a stub function
because the parser doesn't resolve them. So most iOS / Android
fragments DO bind cleanly to the syntax-only scaffolds — only the few
that require a running Activity / ViewController (e.g. lifecycle hook
demonstrations that depend on real `super.onCreate(savedInstanceState)`
behaviour) are listed here. See the per-snippet rationale on each
unbound file for details.
