---
id: python-server-sdk/scaffolds/python-syntax-only
sdk: python-server-sdk
kind: scaffold
lang: python
file: main.py
description: |
  Lite validator that parses the wrappee body without executing it. Use
  for snippets whose runtime context can't be reproduced in the
  validator (fork-based APIs, uWSGI integration, anything requiring
  container orchestration). Catches SyntaxError, indentation bugs, and
  malformed strings — but NOT ImportError or AttributeError on API
  calls. Pair this with full execution validation elsewhere when the
  runtime is available.

  Implementation: the wrappee body is embedded inside a single-quoted
  raw triple-string. A wrappee body containing a literal `'''` would
  break this scaffold; none of the docs snippets we ship today do, but
  document the constraint here in case a future port hits it.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, parsed by ast.parse.
validation:
  runtime: python
  entrypoint: main.py
  # ast.parse is in the stdlib — no third-party deps needed.
---

```python
import ast
import sys

source = r'''
{{ body }}
'''

try:
    ast.parse(source)
except SyntaxError as e:
    print(f"SyntaxError on wrappee body: {e}", file=sys.stderr)
    sys.exit(1)

# The validator harness watches for the EXAM-HELLO success line; emit it
# on a successful parse so a syntax-clean snippet shows as a passing
# validation run.
print("feature flag evaluates to true")
```
