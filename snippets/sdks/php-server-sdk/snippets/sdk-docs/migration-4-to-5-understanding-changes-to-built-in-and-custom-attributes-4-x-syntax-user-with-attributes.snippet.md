---
id: php-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-built-in-and-custom-attributes-4-x-syntax-user-with-attributes
sdk: php-server-sdk
kind: reference
lang: php
description: "4.x syntax, user with attributes in section \"Understanding changes to built-in and custom attributes\""
---

```php
$user = (new LDUserBuilder("example-user-key"))
    ->name("Sandy")
    ->email("sandy@example.com")
    ->custom(["groups" => ["Acme", "Global Health Services"]])
    ->build();
```
