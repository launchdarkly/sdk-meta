---
id: php-server-sdk/sdk-docs/features/storing-data/index/persistent-store
sdk: php-server-sdk
kind: reference
lang: php
description: Persistent feature store (feature requester) configuration example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", [
    'feature_requester' => LaunchDarkly\Integrations\NameOfDatabase::featureRequester()
]);
```
