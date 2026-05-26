---
id: python-server-sdk/ai-configs-agents/strands-full-example
sdk: python-server-sdk
kind: full-example
lang: python
file: python-server-sdk/ai-configs-agents/strands/full-example.txt
description: End-to-end AWS Strands agent (over the OpenAI-compatible API) wired through a LaunchDarkly AI Config agent_config(). Mirrors the "Show full example code" sample in the AI Configs agent quickstart.
---

```python
import ldclient
from ldclient.config import Config
from ldai import LDAIClient
from ldai import AIAgentConfig, AIAgentConfigDefault
from strands import Agent
from strands.models.openai import OpenAIModel

ldclient.set_config(Config("{{sdkkey}}"))
aiclient = LDAIClient(ldclient.get())
context = ldclient.Context.create("user-123")

config = aiclient.agent_config('{{configKey}}', context, AIAgentConfigDefault())
tracker = config.create_tracker()


def handle_agent_call_strands(
    config: AIAgentConfig,
    user_input: str,
) -> str:
    model = config.model.name if config.model else "gpt-5"
    params = config.model.parameters if config.model else {}

    openai_connector = OpenAIModel(
        model_id=model,
        params=params if params else {},
    )

    agent = Agent(system_prompt=config.instructions, model=openai_connector, callback_handler=None)
    response = agent(user_input)
    return str(response)
```
