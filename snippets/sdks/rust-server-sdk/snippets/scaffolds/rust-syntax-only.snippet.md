---
id: rust-server-sdk/scaffolds/rust-syntax-only
sdk: rust-server-sdk
kind: scaffold
lang: rust
file: src/main.rs
description: |
  Parse-only validator for Rust server SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: rust
  entrypoint: src/main.rs
---

```rust
// Pull SDK types into scope so doc-fragment bodies referencing
// `ContextBuilder`, `MultiContextBuilder`, `Reference`, `Client`, etc.
// resolve at compile time without requiring each fragment to repeat
// the imports.
#[allow(unused_imports)]
use launchdarkly_server_sdk::{
    ApplicationInfo, AttributeValue, Client, ConfigBuilder, Context, ContextBuilder,
    MultiContextBuilder, Reason, Reference, ServiceEndpointsBuilder,
    MigratorBuilder, ExecutionOrder, TestData, FlagBuilder,
};
#[allow(unused_imports)]
use std::sync::Arc;
#[allow(unused_imports)]
use futures::future::FutureExt;
#[allow(unused_imports)]
use std::collections::HashMap;

// Placeholder constants the docs reference directly (`YOUR_SDK_KEY`)
// rather than substituting at render time. Declaring them here lets
// fragments like `ConfigBuilder::new(&YOUR_SDK_KEY)` resolve at
// parse time without per-snippet `placeholders:` plumbing.
#[allow(non_upper_case_globals, dead_code)]
const YOUR_SDK_KEY: &str = "";
#[allow(non_upper_case_globals, dead_code)]
const YOUR_MOBILE_KEY: &str = "";
#[allow(non_upper_case_globals, dead_code)]
const YOUR_CLIENT_SIDE_ID: &str = "";
// Some config fragments reference a lowercase `sdk_key` binding rather
// than the YOUR_SDK_KEY placeholder; provide it so they resolve.
#[allow(non_upper_case_globals, dead_code)]
const sdk_key: &str = "";

// Stub for the pre-1.0 (beta) `User` API surface — removed at 1.0
// in favor of Context. Doc fragments under
// `implementation-v1-understanding-*-beta-syntax-*` and
// `implementation-v1-working-with-*-beta-syntax-*` reference the old
// API for side-by-side comparison with the 1.0 equivalents, so the
// syntax-only validator needs a parseable stub to compile them.
//
// Builder methods take `&mut self -> &mut Self` so the docs' multi-line
// shape (`let mut b = User::with_key(...); b.first_name("Sandy");
// b.last_name("Smith"); ... let user = b.build();`) compiles without
// move-after-use errors that would come from a `self -> Self` shape.
// `build(&self) -> Self` lets the final `.build()` produce a User
// without consuming the builder.
#[allow(dead_code, non_camel_case_types)]
struct User;
#[allow(dead_code)]
impl User {
    fn with_key(_key: &str) -> Self { Self }
    fn first_name(&mut self, _v: &str) -> &mut Self { self }
    fn last_name(&mut self, _v: &str) -> &mut Self { self }
    fn email(&mut self, _v: &str) -> &mut Self { self }
    fn custom<T>(&mut self, _v: T) -> &mut Self { self }
    fn build(&self) -> Self { Self }
}

// Stub `hashmap!` macro covering the doc fragments that build a
// `HashMap` for the beta `User::custom(...)` call. Real code reaches
// for `maplit::hashmap!`; declaring the macro inline keeps the
// scaffold free of an extra crate dependency.
#[allow(unused_macros)]
macro_rules! hashmap {
    ($($k:expr => $v:expr),* $(,)?) => {{
        let mut m = ::std::collections::HashMap::new();
        $( m.insert($k, $v); )*
        m
    }};
}

#[allow(dead_code, unused, unused_variables, unused_must_use)]
async fn _wrappee() -> Result<(), Box<dyn std::error::Error>> {
    let client: Client = unimplemented!();
    let context = ContextBuilder::new("stub").build()?;
    // Test-data fragments reference a `td` the docs assume an earlier
    // `TestData::new()` binding created.
    let td = TestData::new();
{{ body }}
    Ok(())
}

#[tokio::main]
async fn main() {
    println!("feature flag evaluates to true");
}
```
