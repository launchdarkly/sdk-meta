---
id: php-server-sdk/sdk-docs/features/storing-data/redis/redis-predis-v63
sdk: php-server-sdk
kind: reference
lang: php
description: Redis (Predis) feature requester configuration example for PHP SDK v6.3 and earlier.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$fr = LaunchDarkly\Integrations\Redis::featureRequester([
    'redis_host' => 'my-redis',
    'redis_port' => 6379,
    'redis_prefix' => 'my-key-prefix'
]);
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", [
    'feature_requester' => $fr
]);
```
