---
id: php-server-sdk/sdk-docs/features/logging/logging
sdk: php-server-sdk
kind: reference
lang: php
description: Custom Monolog logger configuration example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", ["logger" => new Logger("LaunchDarkly", [new ErrorLogHandler(0, Level::Debug)])]);
```
