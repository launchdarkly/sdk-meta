---
id: js-client-sdk/user-feedback/send-feedback-js
sdk: js-client-sdk
kind: reference
lang: javascript
description: sendFeedback helper (JavaScript) with the optional session replay lines commented out.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
// The session replay plugin is optional, but recommended
// import { LDRecord } from "@launchdarkly/session-replay";

export function sendFeedback(
  client,
  flagKey,
  feedback,
  sentiment,
  prompt,
  customProperties,
) {
  const feedbackData = {
    feedback_answer: feedback,
    flag_key: flagKey,
    sentiment: sentiment ?? "neutral",
    custom_properties: customProperties,
  };
  // const sessionID = LDRecord?.getSession()?.sessionSecureID;
  // if (sessionID) {
  //   feedbackData.o11y_session_id = sessionID;
  // }
  if (prompt) {
    feedbackData.feedback_prompt = prompt;
  }

  client.track("$ld:feedback", feedbackData);
  client.flush();
}
```
