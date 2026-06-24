---
id: ruby-server-sdk/scaffolds/ruby-syntax-only-block
sdk: ruby-server-sdk
kind: scaffold
lang: ruby
file: main.rb
description: |
  Block-scope sibling of `ruby-syntax-only`. That scaffold splices
  the body inside `def _wrappee`, which breaks for fragments that
  define a class — Ruby rejects `class` inside a method body at parse
  time. This variant splices the body inside a never-called proc:
  class definitions are legal inside blocks, and the proc body is
  compiled at load without being executed, so references to an
  ambient `client` / `context` never run.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, spliced inside a never-called proc.
validation:
  runtime: ruby
  entrypoint: main.rb
---

```ruby
_wrappee = proc do
{{ body }}
end

puts 'feature flag evaluates to true'
```
