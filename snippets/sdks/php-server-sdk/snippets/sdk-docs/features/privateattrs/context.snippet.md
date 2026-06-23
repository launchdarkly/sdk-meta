---
id: php-server-sdk/sdk-docs/features/privateattrs/context
sdk: php-server-sdk
kind: reference
lang: php
description: Marking context attributes private with the context builder for PHP SDK v5.0.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$context = LDContext::builder('example-context-key')
    ->set('email', 'sandy@example.com')
    ->private('email')
    ->build();
```
