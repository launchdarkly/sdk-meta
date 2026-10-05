---
id: js-client-sdk/user-feedback/send-feedback-session-replay-ts
sdk: js-client-sdk
kind: reference
lang: typescript
description: sendFeedback helper (TypeScript) that attaches the session replay session ID.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```ts
// The session replay plugin is optional, but recommended
import { LDRecord } from "@launchdarkly/session-replay";
import type { LDClient } from "@launchdarkly/js-client-sdk";

export type LDFeedbackData = {
  feedback_answer: string;
  flag_key: string;
  sentiment?: LDFeedbackSentiment;
  feedback_prompt?: string;
  o11y_session_id?: string;
  custom_properties?: Record<string, any>;
};

export type LDFeedbackSentiment = "positive" | "neutral" | "negative";

export function sendFeedback(
  client: LDClient,
  flagKey: string,
  feedback: string,
  sentiment?: LDFeedbackSentiment,
  prompt?: string,
  customProperties?: Record<string, any>,
) {
  const feedbackData: LDFeedbackData = {
    feedback_answer: feedback,
    flag_key: flagKey,
    sentiment: sentiment ?? "neutral",
    custom_properties: customProperties,
  };
  const sessionID = LDRecord?.getSession()?.sessionSecureID;
  if (sessionID) {
    feedbackData.o11y_session_id = sessionID;
  }
  if (prompt) {
    feedbackData.feedback_prompt = prompt;
  }

  client.track("$ld:feedback", feedbackData);
  client.flush();
}
```
