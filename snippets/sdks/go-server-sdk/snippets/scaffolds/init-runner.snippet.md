---
id: go-server-sdk/scaffolds/init-runner
sdk: go-server-sdk
kind: scaffold
lang: go
file: wrappee/init.go
description: |
  Runs an `init.txt`-style snippet end-to-end against a real LaunchDarkly
  env. Layout:
    - wrappee/init.go (this scaffold) — the snippet body verbatim, with
      the `'YOUR_SDK_KEY'` placeholder substituted at validate time. The
      scaffold's `file:` is intentionally `wrappee/init.go` so the
      runner can reference it as a sub-package.
    - main.go (companion `init-runner-main`) — invokes `go run
      wrappee/init.go`, captures its stdout, asserts the wrappee printed
      `SDK successfully initialized!`, then emits the EXAM-HELLO success
      line the harness greps for. The scaffold sets
      `validation.entrypoint: main.go` so the harness runs the runner,
      not the wrappee.

  Two `package main` files in different directories don't conflict;
  Go module resolution picks up the wrappee's imports through the
  shared `go.mod` the harness initializes at the staging-dir root.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, written into wrappee/init.go.
validation:
  runtime: go
  entrypoint: main.go
  companions:
    - go-server-sdk/scaffolds/init-runner-main
---

```go
{{ body }}
```
