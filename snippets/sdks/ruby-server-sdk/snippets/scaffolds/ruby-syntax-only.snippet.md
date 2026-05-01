---
id: ruby-server-sdk/scaffolds/ruby-syntax-only
sdk: ruby-server-sdk
kind: scaffold
lang: ruby
file: main.rb
description: |
  Parse-only validator for Ruby server SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: ruby
  entrypoint: main.rb
---

```ruby
def _wrappee
{{ body }}
end

puts 'feature flag evaluates to true'
```
