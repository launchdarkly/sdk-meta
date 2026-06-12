---
id: js-client-sdk/sdk-docs/features/qualfeedback/send-feedback
sdk: js-client-sdk
kind: reference
lang: javascript
description: Qualitative feedback helper without session replay for JavaScript.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```js
function sendFeedback(client, flagKey, feedback, sentiment, prompt, slackChannelId) {
    // Create an object
    const feedbackData = {
        feedback_answer: feedback,
        flag_key: flagKey,
        sentiment: sentiment,
    }
    // Ensure the provided sentiment value is valid
    const sentiments = ["positive", "neutral", "negative"]
    if (!sentiments.includes(feedbackData.sentiment)) {
      feedbackData.sentiment = "neutral";
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
