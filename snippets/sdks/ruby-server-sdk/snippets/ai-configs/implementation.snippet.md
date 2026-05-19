---
id: ruby-server-sdk/ai-configs/implementation
sdk: ruby-server-sdk
kind: implementation
lang: ruby
file: ruby-server-sdk/ai-configs/implementation.txt
description: Resolve an AI Config with a fallback for ruby-server-sdk.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
fallback_value = LaunchDarkly::AI::AIConfig.new(
  enabled: true,
  model: LaunchDarkly::AI::ModelConfig.new(
    name: 'my-default-model',
    parameters: { 'temperature' => 0.8 }
  ),
  messages: [{ role: 'system', content: '' }],
  provider: { name: 'my-default-provider' }
)

ai_config = ai_client.config('{{configKey}}', context, fallback_value, { 'example_custom_variable' => 'example_custom_value' })
```
