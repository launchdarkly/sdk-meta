---
id: php-server-sdk/sdk-docs/features/storing-data/redis/redis-phpredis
sdk: php-server-sdk
kind: reference
lang: php
description: Redis (phpredis) feature requester configuration example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$redisClient = new Redis();
$redisClient->connect('my-redis', 6379);

$fr = LaunchDarkly\Integrations\PHPRedis::featureRequester($redisClient, ['prefix' => 'my-key-prefix']);
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", [
    'feature_requester' => $fr
]);
```
