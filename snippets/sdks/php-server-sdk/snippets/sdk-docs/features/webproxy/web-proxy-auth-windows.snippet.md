---
id: php-server-sdk/sdk-docs/features/webproxy/web-proxy-auth-windows
sdk: php-server-sdk
kind: reference
lang: php
description: Authenticated web proxy configuration with the Guzzle event publisher for the PHP SDK on Windows.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", [
    "http_proxy" => "http://username:password@my-proxy-host:8080",
    "event_publisher" => LaunchDarkly\Integrations\Guzzle::eventPublisher()
]);
```
