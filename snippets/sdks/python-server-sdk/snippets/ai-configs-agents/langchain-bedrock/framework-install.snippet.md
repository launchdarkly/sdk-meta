---
id: python-server-sdk/ai-configs-agents/langchain-bedrock-framework-install
sdk: python-server-sdk
kind: framework-install
lang: shell
file: python-server-sdk/ai-configs-agents/langchain-bedrock/framework-install.txt
description: Extra pip install lines for wiring LangChain + AWS Bedrock into an AI Config agent. Appended to the base python-server-sdk/ai-configs/install lines.
---

```shell
pip install langchain langchain-aws boto3
```
