---
id: php-server-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode
sdk: php-server-sdk
kind: reference
lang: php
description: Proxy mode configuration example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY",
    [ "base_uri" => "https://your-relay-proxy.com:8030",
      "events_uri" => "https://your-relay-proxy.com:8030" ]);
```
