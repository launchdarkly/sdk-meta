---
id: python-server-sdk/ai-configs-agents/openai-agents-full-example
sdk: python-server-sdk
kind: full-example
lang: python
file: python-server-sdk/ai-configs-agents/openai-agents/full-example.txt
description: End-to-end OpenAI Agents SDK agent wired through a LaunchDarkly AI Config agent_config(). Mirrors the "Show full example code" sample in the AI Configs agent quickstart.
---

```python
import ldclient
from ldclient.config import Config
from ldai import LDAIClient
from ldai import AIAgentConfig, AIAgentConfigDefault
from agents import Agent
from agents.run import Runner

ldclient.set_config(Config("{{sdkkey}}"))
aiclient = LDAIClient(ldclient.get())
context = ldclient.Context.create("user-123")

config = aiclient.agent_config('{{configKey}}', context, AIAgentConfigDefault())
tracker = config.create_tracker()


async def handle_agent_call_openai(
    name: str,
    config: AIAgentConfig,
    user_input: str,
) -> str:
    model = config.model.get_parameter("name") if config.model else "gpt-5"
    root = Agent(
        name=name,
        instructions=config.instructions,
        handoffs=[],
        tools=[],
        model=model,
    )
    response = await Runner.run(root, user_input)
    return response.final_output
```
