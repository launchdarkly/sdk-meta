---
id: php-server-sdk/sdk-docs/features/bigsegments/big-segments
sdk: php-server-sdk
kind: reference
lang: php
description: Big segments Redis store configuration example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
use LaunchDarkly\Types\BigSegmentsConfig;

$redisClient = new Predis\Client([]);
$logger = new Psr\Log\NullLogger();
$bigSegmentsStore = LaunchDarkly\Integrations\Redis::bigSegmentsStore($redisClient, $logger);
$bigSegmentsConfig = new BigSegmentsConfig(store: $bigSegmentsStore);

$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", ['big_segments' => $bigSegmentsConfig]);
```
