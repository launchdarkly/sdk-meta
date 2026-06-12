---
id: php-server-sdk/sdk-docs/features/otel/tracing-hook-environment-id
sdk: php-server-sdk
kind: reference
lang: php
description: OpenTelemetry tracing hook with an explicit environment ID for the PHP SDK.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
$client = new LDClient('YOUR_SDK_KEY', [
    'hooks' => [
        new TracingHook(new TracingHookOptions(
            environmentId: 'example-client-side-id',
        )),
    ],
]);
```
