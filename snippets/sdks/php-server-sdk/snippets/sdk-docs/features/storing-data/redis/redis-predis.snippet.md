---
id: php-server-sdk/sdk-docs/features/storing-data/redis/redis-predis
sdk: php-server-sdk
kind: reference
lang: php
description: Redis (Predis) feature requester configuration example for PHP SDK v6.4+.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$redisClient = new Predis\Client([
    'host' => 'my-redis',
    'port' => 6379
]);

$fr = LaunchDarkly\Integrations\Redis::featureRequester($redisClient, ['prefix' => 'my-key-prefix']);
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", [
    'feature_requester' => $fr
]);
```
