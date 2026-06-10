---
id: python-server-sdk/sdk-docs/features/evaluation-reasons/print-reason
sdk: python-server-sdk
kind: reference
lang: python
description: Reason-object inspection example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
def print_reason(reason):
  kind = reason["kind"]
  if kind == "OFF":
    print("it's off")
  elif kind == "FALLTHROUGH":
    print("fell through")
  elif kind == "TARGET_MATCH":
    print("targeted")
  elif kind == "RULE_MATCH":
    print("matched rule %d/%s" % (reason["ruleIndex"], reason["ruleId"]))
  elif kind == "PREREQUISITE_FAILED":
    print("prereq failed: %s" % reason["prerequisiteKey"])
  elif kind == "ERROR":
    print("error: %s" % reason["errorKind"])
```
