---
id: php-server-sdk/sdk-docs/features/otel/tracing-hook
sdk: php-server-sdk
kind: reference
lang: php
description: OpenTelemetry tracing hook configuration for the PHP SDK.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
require 'vendor/autoload.php';

use LaunchDarkly\LDClient;
use LaunchDarkly\OpenTelemetry\TracingHook;

$client = new LDClient('YOUR_SDK_KEY', [
    'hooks' => [new TracingHook()],
]);
```
