---
id: php-server-sdk/sdk-docs/features/offlinemode/offline-mode
sdk: php-server-sdk
kind: reference
lang: php
description: Offline mode example for PHP SDK v5.0.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", ["offline" => true]);
$client->variation("any.feature.flag", $context, false); // will always return the default value (false)
```
