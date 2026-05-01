---
id: php-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-private-attributes-5-0-syntax-an-attribute-marked-private-for-one-context
sdk: php-server-sdk
kind: reference
lang: php
description: "5.0 syntax, an attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```php
$context = LDContext::builder("example-context-key")
    ->name("Sandy")
    ->set("email", "sandy@example.com")
    ->private("email")
    ->build();
```
