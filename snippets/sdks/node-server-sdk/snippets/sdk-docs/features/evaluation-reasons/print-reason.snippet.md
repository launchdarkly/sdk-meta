---
id: node-server-sdk/sdk-docs/features/evaluation-reasons/print-reason
sdk: node-server-sdk
kind: reference
lang: javascript
description: Reason-object inspection example for Node.js (server-side).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```javascript
function printReason(reason) {
  switch(reason.kind) {
    case "OFF":
      console.log("it's off");
      break;
    case "FALLTHROUGH":
      console.log("fell through");
      break;
    case "TARGET_MATCH":
      console.log("targeted");
      break;
    case "RULE_MATCH":
      console.log("matched rule " + reason.ruleIndex + ", "  + reason.ruleId);
      break;
    case "PREREQUISITE_FAILED":
      console.log("prereq failed: " + reason.prerequisiteKey);
      break;
    case "ERROR":
      console.log("error: " + reason.errorKind);
      break;
  }
}
```
