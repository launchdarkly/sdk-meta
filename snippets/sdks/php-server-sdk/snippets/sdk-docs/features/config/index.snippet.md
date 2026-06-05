---
id: php-server-sdk/sdk-docs/features/config/index
sdk: php-server-sdk
kind: reference
lang: php
description: SDK configuration example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", ["cache" => $cacheStorage, "connect_timeout" => 3]);
```
