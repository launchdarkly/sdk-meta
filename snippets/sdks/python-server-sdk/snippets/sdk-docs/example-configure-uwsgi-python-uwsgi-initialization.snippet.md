---
id: python-server-sdk/sdk-docs/example-configure-uwsgi-python-uwsgi-initialization
sdk: python-server-sdk
kind: reference
lang: python
description: "Python uWSGI initialization in section \"Example: Configure uWSGI\""
validation:
  # uWSGI isn't installable in the validator container; fall back to
  # parse-only. Same constraint as the existing uwsgi-init binding.
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
