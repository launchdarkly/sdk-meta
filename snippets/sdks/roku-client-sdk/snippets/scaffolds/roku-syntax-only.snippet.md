---
id: roku-client-sdk/scaffolds/roku-syntax-only
sdk: roku-client-sdk
kind: scaffold
lang: brightscript
file: main.brs
description: |
  Parse-only validator for Roku BrightScript client SDK doc fragments. No Docker harness yet — scaffold present for future runtime wiring.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: brightscript
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
