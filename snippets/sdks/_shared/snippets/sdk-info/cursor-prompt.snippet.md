---
id: cursor-prompt
kind: reference
lang: text
file: prompt.txt
description: Cursor SDK setup prompt template (uses {{SDK_NAME}}, {{SDK_DOCS_URL}}, {{SDK_EVENT_DOC_URL}} placeholders consumed by gonfalon at runtime).
---

```text
Set up the LaunchDarkly {{SDK_NAME}} SDK in this project by following the instructions here: {{SDK_DOCS_URL}}.

Scope:
- Focus only on validating whether the {{SDK_NAME}} SDK is installed, and initializing it correctly with an environment variable.
- Do not create test files, example projects, or documentation files.
- Ensure initialization makes use of the SDK's `track` functionality to confirm connectivity.

MCP check (required):
  - Look for an MCP server named "LaunchDarkly" that runs @launchdarkly/mcp-server.
    * Cursor: read ~/.cursor/mcp.json
  - MCP is "installed" if the config includes:
      {
        "mcpServers": {
          "LaunchDarkly": {
            "command": "npx",
            "args": ["-y", "--package", "@launchdarkly/mcp-server", "--", "mcp", "start", "--api-key", "<token-or-env-var>"]
          }
        }
      }
  - If not found, offer to add it (prefer an ENV var for the API key).

If MCP is installed and enabled:
  - Use its tools to verify or accelerate setup (read-only by default):
    * List feature flags to confirm API access.
    * (Optional, ask first) Create a temporary "hello-world-sdk-check" flag to validate SDK wiring.
  - If access is restricted, continue SDK setup but skip write operations.

SDK checks:
  1. Is the {{SDK_NAME}} SDK package installed in this repo?
     - If not installed, install it.
  2. Is the SDK already initialized and used?
     - Look for an initialization call.
     - Ensure at least one flag evaluation and a `track` call are present in application code.
     - Verify the SDK key is not hard-coded.

Branch logic:
- If the SDK is installed AND initialized (with flag evaluation + track):
    → Ask user if they'd like to validate configuration (evaluate known flag or the temporary "hello-world-sdk-check" flag).
- If the SDK is installed BUT not initialized:
    → Follow {{SDK_DOCS_URL}} to add initialization, a first flag evaluation, and a `track` call.
    → Follow {{SDK_EVENT_DOC_URL}} to add a custom event called 'source' with the value "cursor" to validate configuration.
- If the SDK is NOT installed:
    → Install it first, then follow {{SDK_DOCS_URL}} for initialization, flag evaluation, and `track` call.
    → Follow {{SDK_EVENT_DOC_URL}} to add a custom event called 'source' with the value "cursor" to validate configuration.

Examples of `track` calls (to guide code insertion):

- **JavaScript (client-side):**
  ```javascript
  const ldClient = useLDClient();
  ldClient.track(PROCESS.env.LAUNCHDARKLY_SDK_KEY, { source: "cursor" });
  ```
- **Python (server-side):**
  ```python
  import os
  from ldclient import Context
  from ldclient.config import Config

  ldclient.set_config(Config(os.getenv("LAUNCHDARKLY_SDK_KEY")))
  context = (Context.builder("context-key-123abc").name("Sandy").build())
  data = {"source": "cursor"}

  ldclient.get().track(os.getenv("LAUNCHDARKLY_SDK_KEY"), context, data)
  ```
- **Java (server-side):**
  ```java
  import com.launchdarkly.sdk.*;
  import com.launchdarkly.sdk.server.*;

  final LDContext context = LDContext.builder("example-context-key")
  .name("Sandy")
  .build();

  final LDConfig config = new LDConfig.Builder().build();

  final LDValue data = LDValue.buildObject().put("source", "cursor").build();
  final LDClient ldClient = new LDClient(System.getenv("LAUNCHDARKLY_SDK_KEY"));
  ldClient.trackData(System.getenv("LAUNCHDARKLY_SDK_KEY"), context, data);
  ```
- **.NET (Client-side):**
  ```csharp
  using LaunchDarkly.Sdk;
  using LaunchDarkly.Sdk.Client;

  Context context = Context.New("context-key-123abc");

  var timeSpan = TimeSpan.FromSeconds(5);
  client = await LdClient.InitAsync("mob-key", ConfigurationBuilder.AutoEnvAttributes.Enabled, context, timeSpan);

  if (client.Initialized)
  {
      var data = LdValue.BuildObject().Add("source", "cursor").Build();
      client.Track(builder.Configuration["LAUNCHDARKLY_SDK_KEY"], data);
      Console.WriteLine("SDK successfully initialized!");
  }
  ```
- **Node.js (server-side):**
  ```javascript
  import * as LaunchDarkly from '@launchdarkly/node-server-sdk';

  const context = {
    kind: 'user',
    key: 'user-key-123abcde',
    email: 'biz@face.dev',
  };

  client.once('ready', function () {
    client.track(PROCESS.env.LAUNCHDARKLY_SDK_KEY, context, { source: "cursor" });
    console.log('SDK successfully initialized!');
  });
  ```
- **React (client-side):**
  ```javascript
  import { LDProvider, useLDClient } from 'launchdarkly-react-client-sdk';
  import { useEffect } from 'react';

  function App() {
    const ldClient = useLDClient();

    useEffect(() => {
      ldClient?.track(PROCESS.env.LAUNCHDARKLY_SDK_KEY, { source: "cursor" });
    }, [ldClient]);

    ...
  }
  ```

SDK key handling:
- Always use an environment variable placeholder appropriate to the language:
    - JavaScript/TypeScript: process.env.LAUNCHDARKLY_SDK_KEY
    - Python: os.getenv("LAUNCHDARKLY_SDK_KEY")
    - Java: System.getenv("LAUNCHDARKLY_SDK_KEY")
    - .NET: builder.Configuration["LAUNCHDARKLY_SDK_KEY"]
    - Go: os.Getenv("LAUNCHDARKLY_SDK_KEY")
- Never hard-code secrets. Show placeholders only.

Validation:
- Run the app and confirm a flag evaluation and `track` event occur.
- If MCP is available:
    * Confirm the flag exists (list/get).
    * (Optional, ask first) Toggle the flag to verify updates are received.
- Report success or provide next steps on error.
```
