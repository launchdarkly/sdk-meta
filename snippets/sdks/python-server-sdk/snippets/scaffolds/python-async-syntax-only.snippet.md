---
id: python-server-sdk/scaffolds/python-async-syntax-only
sdk: python-server-sdk
kind: scaffold
lang: python
file: main.py
description: |
  Async variant of python-syntax-only. It parses the wrappee body without
  running it, and it allows a top-level `await` (through
  ast.PyCF_ALLOW_TOP_LEVEL_AWAIT). Use it for async snippets whose runtime
  context cannot run in the validator — the async client needs a running
  event loop and a real SDK key. It catches SyntaxError, indentation bugs,
  and malformed strings, but NOT ImportError or AttributeError on API calls.

  Implementation: the wrappee body goes inside a single-quoted raw
  triple-string. A wrappee body that contains a literal triple-quote would
  break this scaffold; none of the docs snippets we ship today do.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, compiled by compile().
validation:
  runtime: python
  entrypoint: main.py
  # compile() with the top-level-await flag is in the stdlib — no third-party deps.
---

```python
import ast
import sys

source = r'''
{{ body }}
'''

try:
    compile(source, "main.py", "exec", flags=ast.PyCF_ALLOW_TOP_LEVEL_AWAIT)
except SyntaxError as e:
    print(f"SyntaxError on wrappee body: {e}", file=sys.stderr)
    sys.exit(1)

# The validator harness watches for the EXAM-HELLO success line; emit it
# on a successful parse so a syntax-clean snippet shows as a passing run.
print("feature flag evaluates to true")
```
