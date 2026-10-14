---
id: php-server-sdk/sdk-docs/features/webproxy/web-proxy-auth
sdk: php-server-sdk
kind: reference
lang: php
description: Web proxy configuration with authentication for the PHP SDK.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", [
    "http_proxy" => "http://username:password@my-proxy-host:8080"
]);
```
