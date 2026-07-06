---
id: php-server-sdk/sdk-docs/openfeature/evaluate-a-context
sdk: php-server-sdk
kind: reference
lang: php
file: php-server-sdk/sdk-docs/openfeature/evaluate-a-context.php
description: "PHP OpenFeature provider in section \"Evaluate a context\""
validation:
  scaffold: php-server-sdk/scaffolds/openfeature-php-runner
---

```php
$flagValue = $client->getBooleanValue("example-flag-key", false, $context);
```
