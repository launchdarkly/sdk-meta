---
id: php-server-sdk/sdk-docs/openfeature/initialize-the-provider
sdk: php-server-sdk
kind: reference
lang: php
file: php-server-sdk/sdk-docs/openfeature/initialize-the-provider.php
description: "PHP OpenFeature provider in section \"Initialize the provider\""
validation:
  scaffold: php-server-sdk/scaffolds/openfeature-php-init-runner
  placeholders:
    YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
---

```php
$config = [];
$provider = new LaunchDarkly\OpenFeature\Provider("YOUR_SDK_KEY", $config);
$api = OpenFeatureAPI::getInstance();
$api->setProvider($provider);

$client = $api->getClient("hello-client", 1);
```
