---
id: roku-client-sdk/scaffolds/roku-syntax-only-toplevel
sdk: roku-client-sdk
kind: scaffold
lang: brightscript
file: main.brs
description: |
  File-scope variant of `roku-syntax-only` for Roku BrightScript doc
  fragments that are themselves top-level declarations (named
  `function` / `sub` definitions). BrightScript does not allow a
  named function to be declared inside another function, so those
  fragments cannot route through the `sub _Wrappee()` scaffold.

  Routes through the same `roku-client` Docker validator (the
  RokuCommunity `brighterscript` Lexer + Parser over the staged
  `.brs` file). No Roku device, no real LD env -- nothing in the body
  is ever invoked, and `sub Main()` prints the EXAM-HELLO sentinel
  unconditionally.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, spliced at file scope.
validation:
  runtime: roku-client
  entrypoint: main.brs
---

```brightscript
sub Main()
    print "feature flag evaluates to true"
end sub

{{ body }}
```
