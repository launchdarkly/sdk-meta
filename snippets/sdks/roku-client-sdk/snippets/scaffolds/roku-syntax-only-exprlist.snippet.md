---
id: roku-client-sdk/scaffolds/roku-syntax-only-exprlist
sdk: roku-client-sdk
kind: scaffold
lang: brightscript
file: main.brs
description: |
  Expression-list variant of `roku-syntax-only` for Roku BrightScript
  doc fragments that are a list of expressions, one per line (e.g.
  the supported-log-levels reference). BrightScript forbids a bare
  property access as an expression statement, so such fragments fail
  the `sub _Wrappee()` scaffold; splicing them as the elements of an
  array literal gives every line a legal expression context while
  keeping the body verbatim.

  Routes through the same `roku-client` Docker validator (the
  RokuCommunity `brighterscript` Lexer + Parser). No Roku device, no
  real LD env -- `_Wrappee` is never invoked, and `sub Main()` prints
  the EXAM-HELLO sentinel unconditionally.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, spliced as array-literal elements.
validation:
  runtime: roku-client
  entrypoint: main.brs
---

```brightscript
sub Main()
    print "feature flag evaluates to true"
end sub

sub _Wrappee()
    _expressions = [
{{ body }}
    ]
end sub
```
