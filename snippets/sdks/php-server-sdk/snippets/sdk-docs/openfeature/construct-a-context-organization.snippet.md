---
id: php-server-sdk/sdk-docs/openfeature/construct-a-context-organization
sdk: php-server-sdk
kind: reference
lang: php
file: php-server-sdk/sdk-docs/openfeature/construct-a-context-organization.php
description: "PHP OpenFeature provider in section \"Construct a context\" (organization)"
validation:
  scaffold: php-server-sdk/scaffolds/openfeature-php-runner
---

```php
$attributes = new Attributes(["kind" => "organization"]);
$context = new EvaluationContext("example-user-key", $attributes);
```
