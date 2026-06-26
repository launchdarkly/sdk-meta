---
id: ruby-server-sdk/sdk-docs/features/aimetrics/bedrock-helper
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Helper function that maps config messages to Bedrock Converse arguments for the Ruby AI SDK.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
# The map_converse_arguments helper function transforms the messages
# to fit the Bedrock SDK.

def map_converse_arguments(model_id, messages)
  args = {
    model_id: model_id,
  }

  mapped_messages = []
  user_messages = messages.select { |msg| msg.role == 'user' }
  mapped_messages << { role: 'user', content: user_messages.map { |msg| { text: msg.content } } } unless user_messages.empty?

  assistant_messages = messages.select { |msg| msg.role == 'assistant' }
  mapped_messages << { role: 'assistant', content: assistant_messages.map { |msg| { text: msg.content } } } unless assistant_messages.empty?
  args[:messages] = mapped_messages unless mapped_messages.empty?

  system_messages = messages.select { |msg| msg.role == 'system' }
  args[:system] = system_messages.map { |msg| { text: msg.content } } unless system_messages.empty?

  args
end
```
