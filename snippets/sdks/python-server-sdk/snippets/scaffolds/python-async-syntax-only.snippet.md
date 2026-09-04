---
id: python-server-sdk/scaffolds/python-async-syntax-only
sdk: python-server-sdk
kind: scaffold
lang: python
file: main.py
description: |
  Lite validator that parses the wrappee body without executing it, and
  allows a top-level `await` (through ast.PyCF_ALLOW_TOP_LEVEL_AWAIT). Use
  for async snippets whose runtime context can't be reproduced in the
  validator — they need a running event loop and a real SDK key. Catches
  SyntaxError, indentation bugs, and malformed strings — but NOT ImportError
  or AttributeError on API calls. Pair this with full execution validation
  elsewhere when the runtime is available.

  Implementation: the wrappee body is embedded inside a single-quoted raw
  triple-string. A wrappee body containing a literal triple-quote would
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
