---
id: php-server-sdk/sdk-docs/features/anonymous/anonymous
sdk: php-server-sdk
kind: reference
lang: php
description: Anonymous context example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
$context = LDContext::builder("example-context-key")->anonymous(true)->build();
```
