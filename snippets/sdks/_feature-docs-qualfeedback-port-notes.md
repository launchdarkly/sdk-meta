# Port notes: /sdk/features/user-feedback (qual-feedback)

Source: `ld-docs-private` `fern/topics/sdk/features/qual-feedback.mdx`.
6 code blocks extracted into `sdk-docs/features/qualfeedback/` snippets.
The page is JavaScript-SDK-only, so all six live under `js-client-sdk`;
the two feedback-widget blocks are React component samples and bind
cross-SDK to `react-client-sdk/scaffolds/react-syntax-only`. All six are
bound to validators.

## Content corrections

Fixes applied where the published sample could not work as written.
Everything else is verbatim from the MDX.

- **JavaScript feedback helper, with and without session replay**
  (`js-client-sdk/.../send-feedback-session-replay`,
  `js-client-sdk/.../send-feedback`): two identical bugs in both
  blocks. (1) `if !sentiments.includes(...) {` — JavaScript requires
  parentheses around the condition, so the block was a syntax error;
  rewritten as `if (!sentiments.includes(...)) {`. (2) The fallback
  branch assigned `feedback.sentiment = "neutral"` — `feedback` is the
  string answer parameter, so the normalization never reached the event
  payload; the intended target is `feedbackData.sentiment`.

## Validation routing added in this port

- The four `sendFeedback` helper blocks bind to the existing
  `js-client-sdk/scaffolds/js-syntax-only` scaffold. The js-client
  validator image already carries `@launchdarkly/session-replay`, so
  the `LDRecord` import bundles as-is.
- The two React widget blocks contain JSX, which the js-client
  validator's `src/app.ts` staging cannot parse; they bind to
  `react-client-sdk/scaffolds/react-syntax-only` (scaffold resolution
  is global, so a `js-client-sdk` snippet can wrap itself in another
  SDK's scaffold).
- react-client validator image: added a stub `src/sendFeedback.ts`
  module (exports `sendFeedback` and the `LDFeedbackSentiment` type)
  so the widget blocks' `import ... from './sendFeedback'` resolves at
  bundle time. Mirrors the existing `src/environments.ts` stub.

## Known non-binds

None.
