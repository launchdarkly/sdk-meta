---
id: roku-client-sdk/scaffolds/roku-syntax-only
sdk: roku-client-sdk
kind: scaffold
lang: brightscript
file: main.brs
description: |
  Parse-only validator for Roku BrightScript client SDK doc
  fragments. Routes through the `roku-client` Docker validator
  which runs the RokuCommunity `brighterscript` Lexer + Parser
  over the staged `.brs` file. No Roku device, no real LD env —
  the wrappee body sits inside `sub _Wrappee()` (never invoked)
  and `sub Main()` prints the EXAM-HELLO sentinel unconditionally.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: roku-client
  entrypoint: main.brs
---

```brightscript
sub Main()
    print "feature flag evaluates to true"
end sub

sub _Wrappee()
{{ body }}
end sub
```
