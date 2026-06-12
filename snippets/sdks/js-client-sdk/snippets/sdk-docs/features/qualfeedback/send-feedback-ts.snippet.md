---
id: js-client-sdk/sdk-docs/features/qualfeedback/send-feedback-ts
sdk: js-client-sdk
kind: reference
lang: typescript
description: Qualitative feedback helper without session replay for TypeScript.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```ts
// Create an object
type LDFeedbackData = {
    feedback_answer: string;
    flag_key: string;
    sentiment?: LDFeedbackSentiment;
    feedback_prompt?: string;
    o11y_session_id?: string;
}

// Ensure the provided sentiment value is valid
type LDFeedbackSentiment = "positive" | "neutral" | "negative";

function sendFeedback(
    client: LDClient,
    flagKey: string,
    feedback: string,
    sentiment?: LDFeedbackSentiment,
    prompt?: string,
    slackChannelId?: string,
) {
    const feedbackData: LDFeedbackData = {
        feedback_answer: feedback,
        flag_key: flagKey,
        sentiment: sentiment ?? "neutral",
    }
    // The wording of the question or form prompt that generated the feedback
    if (prompt) {
        feedbackData.feedback_prompt = prompt;
    }

    // Send the event back to LaunchDarkly
    client.track('$ld:feedback', feedbackData);
    client.flush();
}
```
