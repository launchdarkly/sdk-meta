---
id: php-server-sdk/sdk-docs/evaluate-a-context-php
sdk: php-server-sdk
kind: reference
lang: php
description: "PHP in section \"Evaluate a context\""
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
$context = LDContext::builder("example-context-key")
  ->name("Sandy")
  ->build();

if ($client->variation("your.flag.key", $context)) {
    // application code to show the feature
} else {
    // the code to run if the feature is off
}
```
