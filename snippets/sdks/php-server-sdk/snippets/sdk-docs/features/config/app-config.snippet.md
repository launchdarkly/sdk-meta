---
id: php-server-sdk/sdk-docs/features/config/app-config
sdk: php-server-sdk
kind: reference
lang: php
description: Application metadata configuration example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$appInfo = (new ApplicationInfo())->withId('authentication-service')->withVersion('1.0.0');
$config = [
  "application_info" => $appInfo
];

$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", $config);
```
