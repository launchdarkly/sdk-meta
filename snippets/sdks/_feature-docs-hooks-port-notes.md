# Port notes: /sdk/features/hooks

Source: `ld-docs-private` `fern/topics/sdk/features/hooks.mdx`.
20 code blocks extracted into `sdk-docs/features/hooks/` snippets
across 11 SDKs. All 20 are bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX (modulo indentation flush to
column 0 for the JavaScript section's indented CodeBlocks).

- **.NET (server) hook definition** (`dotnet-server-sdk/.../define-hook`):
  `Hook` has a single constructor, `Hook(string name)` — no
  parameterless one — so `class ExampleHook : Hook { }` cannot supply
  an implicit constructor and `new ExampleHook()` cannot compile.
  Added `public ExampleHook() : base("example-hook") { }`.
- **Android hook definition** (`android-client-sdk/.../define-hook`):
  same problem — the android `Hook` class only has `Hook(String name)`,
  and the sample instantiates `new ExampleHook("Example hook")` without
  declaring a matching constructor. Added
  `public ExampleHook(String name) { super(name); }`, mirroring the
  Java server sample on the same page.
- **Android add-hook** (`android-client-sdk/.../add-hook`):
  `client.setHooks(hookList)` does not exist on `LDClient` — the only
  post-init registration API is `addHook(Hook)` (since 5.8.0).
  Replaced the call with `client.addHook(exampleHook);`.
- **Python hook definition** (`python-server-sdk/.../define-hook`):
  the "Implement at least one of" comment listed `before_evaluation`
  twice; the second one is now `after_evaluation`, matching the stage
  descriptions that follow.

Known-but-left-verbatim: the Go sample's `func newExampleHook()
exampleHook {}` sketch has no return statement and `client, _ = ...`
assigns to an ambient `client` the fragment never declares. Both are
consistent with the page's sketch style (ambient names are pervasive
in these fragments), parse cleanly, and were left untouched.

## Validation routing added in this port

- `go-server-sdk/scaffolds/go-syntax-only-raw` (+`-main` companion) —
  `go-syntax-only` embeds the body in a backtick-delimited Go raw
  string, and the hooks body's comments contain backticks
  (`` `BeforeEvaluation` ``), which terminate the literal. The new
  pair stages the body verbatim as `fragment.txt` and the companion
  `main.go` reads it from disk; split/parse semantics are unchanged.
- `node-server-sdk/scaffolds/node-syntax-only-toplevel` (+`-checker`
  companion) — same backtick problem with `node-syntax-only`'s
  template-literal embedding, plus the body's `implements
  integrations.Hook` clause is TypeScript that `node --check` rejects.
  The pair stages the body verbatim as `fragment.ts`; the checker
  erases the simple TS surface (the `node-syntax-only` regexes plus an
  `implements X {` eraser) and runs `node --check` on the result.
- `ruby-server-sdk/scaffolds/ruby-syntax-only-toplevel` — Ruby rejects
  `class` definitions inside a method body, so the def-wrapped
  `ruby-syntax-only` cannot host the hook class. The new scaffold
  embeds the body in a single-quoted heredoc and parses it with
  `RubyVM::AbstractSyntaxTree.parse` (no execution).
- TYPE_LIFT pre-stage rewrite (dotnet-server + android-client
  harnesses) — C# has no local type declarations at all, and Java
  rejects access modifiers on local classes, so "define a class, then
  configure the client with it" fragments cannot compile inside the
  syntax-only scaffolds' wrapper methods. `csharp-syntax-only` and the
  android `java-syntax-only` scaffolds now carry `// TYPE_LIFT_TARGET`
  and `// BODY_BEGIN` / `// BODY_END` markers; a marker-gated Python
  pass in each harness moves brace-balanced type declarations from the
  body region to the target (namespace scope for C#, SnippetActivity
  member scope for android). Bodies without type declarations are
  untouched, so existing bindings behave as before.
- Stub-surface extensions: `java.util.Collections` import on the Java
  server `java-syntax-only` scaffold (the sample calls
  `Collections.singletonList` without importing it — the docs assume
  it); `java.util.List` / `java.util.ArrayList` imports, an
  `LDConfig ldConfig` field, and a file-scope `ExampleHook` stub class
  on the android `java-syntax-only` scaffold (the add-hook fragment
  registers a hook class defined in the preceding example; a
  type-lifted body class shadows the stub).

## Known non-binds

None. The page has no Objective-C block; the iOS Swift snippet is
bound to `swift-syntax-only` and validates on the PR's macOS CI row
(the native harness cannot run locally).
