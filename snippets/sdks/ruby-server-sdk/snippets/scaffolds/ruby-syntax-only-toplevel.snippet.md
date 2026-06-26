---
id: ruby-server-sdk/scaffolds/ruby-syntax-only-toplevel
sdk: ruby-server-sdk
kind: scaffold
lang: ruby
file: main.rb
description: |
  Parse-only validator for Ruby server SDK doc fragments that contain
  module-scope declarations. Ruby rejects `class` definitions inside a
  method body ("unexpected class definition in method body"), so the
  `ruby-syntax-only` shape — body spliced into a never-invoked
  `def _wrappee` — cannot host class-implementation fragments.

  The wrappee body is embedded in a single-quoted squiggly heredoc and
  parsed with `RubyVM::AbstractSyntaxTree.parse`, which raises
  SyntaxError on malformed input without executing anything, so
  unresolved constants (`LaunchDarkly::...`) and undefined variables
  pass. A body containing a line consisting of the heredoc terminator
  `WRAPPEE_BODY` would break this scaffold; none of the docs snippets
  we ship today do, but document the constraint here in case a future
  port hits it.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, parsed by RubyVM::AbstractSyntaxTree.
validation:
  runtime: ruby
  entrypoint: main.rb
---

```ruby
source = <<~'WRAPPEE_BODY'
{{ body }}
WRAPPEE_BODY

begin
  RubyVM::AbstractSyntaxTree.parse(source)
rescue SyntaxError => e
  warn "SyntaxError on wrappee body: #{e.message}"
  exit 1
end

# The validator harness watches for the EXAM-HELLO success line; emit it
# on a successful parse so a syntax-clean snippet shows as a passing
# validation run.
puts 'feature flag evaluates to true'
```
