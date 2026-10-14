---
id: php-server-sdk/sdk-docs/features/webproxy/web-proxy
sdk: php-server-sdk
kind: reference
lang: php
description: Web proxy configuration for the PHP SDK.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", [
    "http_proxy" => "http://my-proxy-host:8080"
]);
```
