---
id: rust-server-sdk/sdk-docs/features/evaluation-reasons/print-reason
sdk: rust-server-sdk
kind: reference
lang: rust
description: Reason-object inspection example for Rust.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only

---

```rust
fn print_reason(reason: Reason) {
    match reason {
        Reason::Off => println!("it's off"),
        Reason::Fallthrough { .. } => println!("fell through"),
        Reason::TargetMatch => println!("targeted"),
        Reason::RuleMatch {
            rule_index,
            rule_id,
            ..
        } => println!("matched rule {}/{}", rule_index, rule_id),
        Reason::PrerequisiteFailed { prerequisite_key } => {
            println!("prereq failed: {}", prerequisite_key)
        }
        Reason::Error { error } => println!("error: {:?}", error),
    };
}
```
