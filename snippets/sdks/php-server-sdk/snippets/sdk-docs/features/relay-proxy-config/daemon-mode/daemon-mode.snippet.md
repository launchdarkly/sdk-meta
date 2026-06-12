---
id: php-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode
sdk: php-server-sdk
kind: reference
lang: php
description: Daemon mode configuration example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY",
    [ 'feature_requester' => LaunchDarkly\Integrations\NameOfDatabase::featureRequester() ]);
```
