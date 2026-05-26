---
id: python-server-sdk/ai-configs-agents/claude-full-example
sdk: python-server-sdk
kind: full-example
lang: python
file: python-server-sdk/ai-configs-agents/claude/full-example.txt
description: End-to-end Claude Agent SDK agent wired through a LaunchDarkly AI Config agent_config(). Mirrors the "Show full example code" sample in the AI Configs agent quickstart.
---

```python
import ldclient
from ldclient.config import Config
from ldai import LDAIClient
from ldai import AIAgentConfig, AIAgentConfigDefault
from claude_agent_sdk import query, ClaudeAgentOptions
from claude_agent_sdk.types import ResultMessage

ldclient.set_config(Config("{{sdkkey}}"))
aiclient = LDAIClient(ldclient.get())
context = ldclient.Context.create("user-123")

config = aiclient.agent_config('{{configKey}}', context, AIAgentConfigDefault())
tracker = config.create_tracker()


async def handle_agent_call_claude(
    config: AIAgentConfig,
    user_input: str,
) -> str:
    model = config.model.name if config.model else "claude-sonnet-4-6-20260217"

    final_message = None
    async for message in query(
        prompt=user_input,
        options=ClaudeAgentOptions(
            system_prompt=config.instructions or "",
            model=model,
        ),
    ):
        final_message = message

    if not isinstance(final_message, ResultMessage):
        raise ValueError(f"Unexpected final message type: {type(final_message)}")

    return final_message.result or ""
```
