---
id: php-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-built-in-and-custom-attributes-5-0-syntax-context-with-object-attributes
sdk: php-server-sdk
kind: reference
lang: php
description: "5.0 syntax, context with object attributes in section \"Understanding changes to built-in and custom attributes\""
---

```php
$context = LDContext::builder("example-context-key")
    ->set("address", ["street" => "Main St", "city" => "Springfield"])
    ->build();
```
