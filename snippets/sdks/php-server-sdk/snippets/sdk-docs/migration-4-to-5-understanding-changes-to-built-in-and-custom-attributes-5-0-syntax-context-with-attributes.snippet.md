---
id: php-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-built-in-and-custom-attributes-5-0-syntax-context-with-attributes
sdk: php-server-sdk
kind: reference
lang: php
description: "5.0 syntax, context with attributes in section \"Understanding changes to built-in and custom attributes\""
---

```php
$context = LDContext::builder("example-context-key")
    ->name("Sandy")
    ->set("email", "sandy@example.com")
    ->set("groups", ["Acme", "Global Health Services"])
    ->build();
```
