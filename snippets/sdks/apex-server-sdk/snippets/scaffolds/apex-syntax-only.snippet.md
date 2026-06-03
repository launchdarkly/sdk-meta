---
id: apex-server-sdk/scaffolds/apex-syntax-only
sdk: apex-server-sdk
kind: scaffold
lang: java
file: snippet.apex
description: |
  Parse-only validator for Apex server SDK doc fragments. Routes
  through the `apex` Docker validator, which runs
  `prettier-plugin-apex`'s `apex-anonymous` parser over the staged
  fragment — the same parser the apex-server-sdk repo's CI uses for
  its `prettier --check "**/*.cls"` step. No Salesforce scratch org,
  no LD env: a clean parse means the fragment is syntactically valid
  Apex (Execute Anonymous form). Doc fragments are bare statement
  blocks, so the scaffold emits the body verbatim — the anonymous
  parser accepts a sequence of statements with no surrounding class.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, parsed as anonymous Apex.
validation:
  runtime: apex
  entrypoint: snippet.apex
---

```java
{{ body }}
```
