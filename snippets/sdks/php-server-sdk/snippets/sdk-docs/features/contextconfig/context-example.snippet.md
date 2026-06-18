---
id: php-server-sdk/sdk-docs/features/contextconfig/context-example
sdk: php-server-sdk
kind: reference
lang: php
description: Context example for PHP SDK v5.0.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$context = LDContext::builder("example-context-key")
    ->set("firstName", "Sandy")
    ->set("lastName", "Smith")
    ->set("email", "sandy@example.com")
    ->set("groups", ["Acme", "Global Health Services"])
    ->build();
```
