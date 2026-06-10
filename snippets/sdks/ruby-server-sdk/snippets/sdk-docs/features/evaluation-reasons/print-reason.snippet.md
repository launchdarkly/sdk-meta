---
id: ruby-server-sdk/sdk-docs/features/evaluation-reasons/print-reason
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Reason-object inspection example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
def print_reason(reason)
  case reason[:kind]
  when "OFF"
    puts "it's off"
  when "FALLTHROUGH"
    puts "fell through"
  when "TARGET_MATCH"
    puts "targeted"
  when "RULE_MATCH"
    puts "matched rule #{reason[:ruleIndex]}/#{reason[:ruleId]}"
  when "PREREQUISITE_FAILED"
    puts "prereq failed: #{reason[:prerequisiteKey]}"
  when "ERROR"
    puts "error: #{reason[:errorKind]}"
  end
end
```
