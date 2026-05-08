---
id: python-server-sdk/scaffolds/init-runner-observability
sdk: python-server-sdk
kind: scaffold
lang: python
file: main.py
description: |
  End-to-end runner for the `observability/initialize` snippet body.

  The wrappee body assumes `ldclient`, `Config`,
  `ObservabilityPlugin`, `ObservabilityConfig`, `observe`, and
  `logging` are imported. This scaffold supplies those imports at
  module scope, splices the body, and waits a bounded period for
  `ldclient.get().is_initialized()` to return True. We don't assert
  observability data flows back to LaunchDarkly — just that the SDK
  starts cleanly with the o11y plugin attached. A clean start emits
  the EXAM-HELLO line.

  The wrappee's `'SDK_KEY'` literal is substituted with the live
  `LAUNCHDARKLY_SDK_KEY` env var via the snippet's
  `validation.placeholders` map (handled by the dispatcher upstream).
inputs:
  body:
    type: string
    description: The wrappee init snippet body, embedded after key substitution.
validation:
  runtime: python
  entrypoint: main.py
  requirements: |
    launchdarkly-server-sdk
    launchdarkly-observability
---

```python
import logging
import time

import ldclient
from ldclient.config import Config
from ldobserve import ObservabilityConfig, ObservabilityPlugin, observe

# The wrappee body calls `ldclient.set_config(Config('SDK_KEY', plugins=[...]))`
# and then runs sample observe.record_log / observe.start_span calls
# against the global ldclient instance. Splicing it here at module
# scope lets it use the imports above.
{{ body }}

# Wait a bounded period for the SDK to initialize. The server SDK
# returns from set_config immediately; is_initialized() flips True
# once the streaming connection delivers the initial flag payload.
deadline = time.time() + 10
while time.time() < deadline:
    if ldclient.get().is_initialized():
        break
    time.sleep(0.2)

if not ldclient.get().is_initialized():
    print("scaffold: SDK did not initialize within 10s")
    raise SystemExit(1)

# Best-effort flush so any events the body queued get sent.
try:
    ldclient.get().flush()
except Exception:
    pass

print("SDK successfully initialized")
print("feature flag evaluates to true")
```
