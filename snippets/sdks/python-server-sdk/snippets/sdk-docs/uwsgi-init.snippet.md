---
id: python-server-sdk/sdk-docs/uwsgi-init
sdk: python-server-sdk
kind: reference
lang: python
description: uWSGI integration that calls postfork() inside @uwsgidecorators.postfork.
validation:
  # uWSGI isn't installable in the validator container; fall back to
  # parse-only. The decorator-based fork lifecycle would need a real
  # uWSGI master process to exercise.
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
import uwsgidecorators

# Instantiate the client before the fork
ldclient.set_config(LDConfig("YOUR_SDK_KEY"))
client = ldclient.get()

@uwsgidecorators.postfork
def post_fork_client_initialization():
    # Reinitialize the client after the fork
    ldclient.get().postfork()

```
