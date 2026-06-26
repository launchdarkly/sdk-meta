---
id: php-server-sdk/sdk-docs/features/storing-data/consul/consul
sdk: php-server-sdk
kind: reference
lang: php
description: Consul feature requester configuration example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$fr = LaunchDarkly\Integrations\Consul::featureRequester([
    'consul_uri' => 'http://my-consul:8100',
    'consul_prefix' => 'my-key-prefix'
]);
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", [
    'feature_requester' => $fr
]);
```
