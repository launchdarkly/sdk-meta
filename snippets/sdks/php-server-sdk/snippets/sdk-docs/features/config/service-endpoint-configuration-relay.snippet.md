---
id: php-server-sdk/sdk-docs/features/config/service-endpoint-configuration-relay
sdk: php-server-sdk
kind: reference
lang: php
description: Service endpoint configuration example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY",
    [ "base_uri" => "https://your-relay-proxy.com:8030",
      "events_uri" => "https://your-relay-proxy.com:8030" ]);
```
