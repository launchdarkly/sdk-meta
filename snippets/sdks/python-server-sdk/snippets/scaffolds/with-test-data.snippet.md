---
id: python-server-sdk/scaffolds/with-test-data
sdk: python-server-sdk
kind: scaffold
lang: python
file: main.py
description: |
  Standard validator harness for Python doc fragments that assume a
  pre-initialized `client` and `context`. Initializes the SDK against a
  TestData source pre-populated with the docs' canonical
  `your.feature.key` flag (returns true), builds a Context, runs the
  wrappee body, then prints the EXAM-HELLO success line so the existing
  Python validator harness recognizes a successful run.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted between client setup and the success-line print.
validation:
  runtime: python
  entrypoint: main.py
  # observability is the optional plugin the install/import docs reference;
  # including it here means snippets that mention `ldobserve` imports can
  # validate cleanly under this scaffold without a separate variant.
  requirements: |
    launchdarkly-server-sdk
    launchdarkly-observability
---

```python
import ldclient
from ldclient import Context
from ldclient.config import Config
from ldclient.integrations.test_data import TestData

# TestData is the SDK's in-process flag store — no network, no real LD env.
# We pre-populate `your.feature.key` (the canonical flag key the docs use
# in their evaluating / variation_detail / all_flags examples) to true, so
# any wrappee that evaluates that flag will get a non-default value.
td = TestData.data_source()
td.update(td.flag("your.feature.key").variation_for_all(True))

config = Config("test-sdk-key", update_processor_class=td, send_events=False)
ldclient.set_config(config)
client = ldclient.get()

# Most docs fragments build their own context; this default exists so the
# fragment can omit the builder when it's not the focus of the snippet.
context = Context.builder("example-context-key").name("Sandy").build()

# --- wrappee body ---
{{ body }}
# --- end wrappee body ---

# Final evaluation drives the EXAM-HELLO success line. Snippets that
# already set `flag_value` will see it overwritten here — that's fine, the
# point is just to confirm the SDK initialized and the API surface
# resolves cleanly. A wrappee with stale imports or removed methods would
# have failed before reaching this line.
flag_value = client.variation("your.feature.key", context, False)
print(f"feature flag evaluates to {flag_value}")
```
