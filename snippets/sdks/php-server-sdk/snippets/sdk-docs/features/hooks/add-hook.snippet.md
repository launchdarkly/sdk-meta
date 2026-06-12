---
id: php-server-sdk/sdk-docs/features/hooks/add-hook
sdk: php-server-sdk
kind: reference
lang: php
description: Adding a hook to an existing client for the PHP SDK.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
$client = new LDClient('YOUR_SDK_KEY');
$client->addHook(new ExampleHook());
```
