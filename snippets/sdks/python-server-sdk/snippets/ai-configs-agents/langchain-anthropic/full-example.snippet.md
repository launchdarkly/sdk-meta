---
id: python-server-sdk/ai-configs-agents/langchain-anthropic-full-example
sdk: python-server-sdk
kind: full-example
lang: python
file: python-server-sdk/ai-configs-agents/langchain-anthropic/full-example.txt
description: End-to-end LangChain + Anthropic agent wired through a LaunchDarkly AI Config agent_config(). Mirrors the "Show full example code" sample in the AI Configs agent quickstart.
---

```python
import ldclient
from ldclient.config import Config
from ldai import LDAIClient
from ldai import AIAgentConfig, AIAgentConfigDefault
from langchain.agents import create_agent
from langchain.messages import HumanMessage

ldclient.set_config(Config("{{sdkkey}}"))
aiclient = LDAIClient(ldclient.get())
context = ldclient.Context.create("user-123")

config = aiclient.agent_config('{{configKey}}', context, AIAgentConfigDefault())
tracker = config.create_tracker()


def handle_agent_call_langchain(
    config: AIAgentConfig,
    user_input: str,
) -> str:
    model = config.model.name if config.model else "anthropic:claude-sonnet-4-6"

    agent = create_agent(
        model=model,
        system_prompt=config.instructions,
    )

    response = agent.invoke({"messages": [HumanMessage(user_input)]})
    return response["messages"][-1].content
```
