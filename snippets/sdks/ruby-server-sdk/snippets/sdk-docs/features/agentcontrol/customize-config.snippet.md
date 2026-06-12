---
id: ruby-server-sdk/sdk-docs/features/agentcontrol/customize-config
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Customize an AgentControl config for Ruby AI.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
key = 'example-config-key'
context = LaunchDarkly::LDContext.create({ key: 'example-user-key', kind: 'user', name: 'Sandy' })
fallback_value = LaunchDarkly::Server::AI::AIConfig.new(enabled: false)
variables = { 'example_custom_variable' => 'example_custom_value' }

ai_config = ai_client.config(key, context, fallback_value, variables)

if ai_config.enabled
  # Send a request to your AI provider using the customized config
else
  # Application path to take when the config is disabled
end
```
