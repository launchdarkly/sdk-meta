---
id: js-client-sdk/user-feedback/smiley-popover-js
sdk: js-client-sdk
kind: reference
lang: javascript
description: React smiley popover feedback component (JavaScript). FLAG_KEY and FEEDBACK_PROMPT are substituted by the consumer.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```js
import { useState } from 'react';
import { sendFeedback } from './sendFeedback';

const HappyFaceIcon = () => (
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
    <g stroke="#008344" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5" clipPath="url(#a)">
      <path
        d="M3 12a9 9 0 1 0 18.001 0A9 9 0 0 0 3 12M9 10h.01M15 10h.01"
        style={{ stroke: '#008344', strokeOpacity: 1 }}
      />
      <path d="M9.5 15a3.5 3.5 0 0 0 5 0" style={{ stroke: '#008344', strokeOpacity: 1 }} />
    </g>
    <defs>
      <clipPath id="a">
        <path fill="#fff" d="M0 0h24v24H0z" style={{ fill: '#fff', fillOpacity: 1 }} />
      </clipPath>
    </defs>
  </svg>
);

const NeutralFaceIcon = () => (
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
    <g stroke="#898e94" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5" clipPath="url(#b)">
      <path
        d="M3 12a9 9 0 1 0 18.001 0A9 9 0 0 0 3 12M9 10h.01M15 10h.01M9 15h6"
        style={{ stroke: '#898e94', strokeOpacity: 1 }}
      />
    </g>
    <defs>
      <clipPath id="b">
        <path fill="#fff" d="M0 0h24v24H0z" style={{ fill: '#fff', fillOpacity: 1 }} />
      </clipPath>
    </defs>
  </svg>
);

const FrownyFaceIcon = () => (
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
    <g stroke="#db2251" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5" clipPath="url(#c)">
      <path
        d="M3 12a9 9 0 1 0 18.001 0A9 9 0 0 0 3 12M9 10h.01M15 10h.01"
        style={{ stroke: '#db2251', strokeOpacity: 1 }}
      />
      <path d="M9.5 15.25a3.5 3.5 0 0 1 5 0" style={{ stroke: '#db2251', strokeOpacity: 1 }} />
    </g>
    <defs>
      <clipPath id="c">
        <path fill="#fff" d="M0 0h24v24H0z" style={{ fill: '#fff', fillOpacity: 1 }} />
      </clipPath>
    </defs>
  </svg>
);

export function SmileyFeedbackPopover({ flagKey = "FLAG_KEY", ldClient }) {
  const [isOpen, setIsOpen] = useState(false);
  const [feedback, setFeedback] = useState("");
  const [sentiment, setSentiment] = useState(undefined);

  const handleSubmit = () => {
    sendFeedback(
      ldClient,
      flagKey,
      feedback,
      sentiment,
      "FEEDBACK_PROMPT"
    );
    setIsOpen(false);
    setFeedback("");
    setSentiment(undefined);
  };

  return (
    <div style={{ position: "relative", display: "inline-block" }}>
      <button
        onClick={() => setIsOpen(!isOpen)}
        style={{
          backgroundColor: "black",
          color: "white",
          border: "none",
          borderRadius: "0.25rem",
          padding: "0.5rem 1rem",
          cursor: "pointer",
        }}
      >
        Give feedback
      </button>
      {isOpen && (
        <div
          style={{
            position: "absolute",
            top: "100%",
            left: 0,
            marginTop: "0.5rem",
            backgroundColor: "white",
            border: "1px solid #ccc",
            borderRadius: "0.25rem",
            padding: "16px",
            boxSizing: "border-box",
            width: "300px",
            zIndex: 1000,
          }}
        >
          <textarea
            name="feedback"
            value={feedback}
            onChange={(e) => setFeedback(e.target.value)}
            placeholder="FEEDBACK_PROMPT"
            rows={3}
            style={{
              width: "100%",
              padding: "0.5rem",
              marginBottom: "0.5rem",
              boxSizing: "border-box",
              border: "1px solid #ccc",
              borderRadius: "0.25rem",
              fontFamily: "sans-serif",
            }}
          />
          <div
            style={{
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
            }}
          >
            <div style={{ display: "flex", gap: "0.25rem" }}>
              <button
                onClick={() =>
                  setSentiment(
                    sentiment === "positive" ? undefined : "positive"
                  )
                }
                aria-label="Happy face"
                style={{
                  cursor: "pointer",
                  background: sentiment === "positive" ? "#eee" : "none",
                  border: "none",
                  borderRadius: "0.25rem",
                  padding: "0.25rem 0.5rem",
                }}
              >
                <HappyFaceIcon />
              </button>
              <button
                onClick={() =>
                  setSentiment(
                    sentiment === "neutral" ? undefined : "neutral"
                  )
                }
                aria-label="Neutral face"
                style={{
                  cursor: "pointer",
                  background: sentiment === "neutral" ? "#eee" : "none",
                  border: "none",
                  borderRadius: "0.25rem",
                  padding: "0.25rem 0.5rem",
                }}
              >
                <NeutralFaceIcon />
              </button>
              <button
                onClick={() =>
                  setSentiment(
                    sentiment === "negative" ? undefined : "negative"
                  )
                }
                aria-label="Frowny face"
                style={{
                  cursor: "pointer",
                  background: sentiment === "negative" ? "#eee" : "none",
                  border: "none",
                  borderRadius: "0.25rem",
                  padding: "0.25rem 0.5rem",
                }}
              >
                <FrownyFaceIcon />
              </button>
            </div>
            <button
              onClick={handleSubmit}
              style={{
                backgroundColor: "black",
                color: "white",
                border: "none",
                borderRadius: "0.25rem",
                padding: "0.5rem 1rem",
                cursor: "pointer",
              }}
            >
              Send
            </button>
          </div>
        </div>
      )}
    </div>
  );
}

// Usage example:
// <SmileyFeedbackPopover ldClient={ldClient} />
```
