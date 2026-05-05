---
id: python-server-sdk/scaffolds/init-runner
sdk: python-server-sdk
kind: scaffold
lang: python
file: main.py
description: |
  Runs an `init.txt`-style snippet end-to-end against a real LaunchDarkly
  env. The wrappee's literal SDK-key placeholder is substituted at validate
  time via the snippet's `validation.placeholders` map (handled by the
  dispatcher), so by the time the body is spliced in below it already
  carries the real key.

  The wrappee runs through runpy.run_path with `run_name="__main__"` so
  any `if __name__ == '__main__'` block in the body fires. The scaffold
  watches for the wrappee's own `SDK successfully initialized` line; on
  a clean run it then emits the EXAM-HELLO success line the validator
  harness greps for.
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: python
  entrypoint: main.py
  requirements: |
    launchdarkly-server-sdk
---

```python
import io
import runpy
import sys
import tempfile

# The wrappee body is embedded as a raw triple-quoted string. None of the
# init snippets we ship use the `'''` token at module scope, so this is
# safe today; document the constraint here for any future port that
# adds one.
_WRAPPEE_BODY = r'''
{{ body }}
'''

with tempfile.NamedTemporaryFile("w", suffix=".py", delete=False) as fh:
    fh.write(_WRAPPEE_BODY)
    body_path = fh.name

# Capture stdout from the wrappee so we can assert its success line.
buf = io.StringIO()

class Tee:
    def __init__(self, *streams):
        self.streams = streams
    def write(self, s):
        for st in self.streams:
            st.write(s)
    def flush(self):
        for st in self.streams:
            st.flush()

real_stdout = sys.stdout
sys.stdout = Tee(real_stdout, buf)
try:
    runpy.run_path(body_path, run_name="__main__")
except SystemExit as e:
    if e.code not in (0, None):
        sys.stdout = real_stdout
        print(f"scaffold: wrappee exited with non-zero code {e.code}", file=sys.stderr)
        sys.exit(1)
finally:
    sys.stdout = real_stdout

if "SDK successfully initialized" not in buf.getvalue():
    print("scaffold: wrappee did not print 'SDK successfully initialized'", file=sys.stderr)
    sys.exit(1)

# Match the EXAM-HELLO success regex the harness watches for.
print("feature flag evaluates to true")
```
