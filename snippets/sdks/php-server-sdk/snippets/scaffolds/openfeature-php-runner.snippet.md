---
id: php-server-sdk/scaffolds/openfeature-php-runner
sdk: php-server-sdk
kind: scaffold
lang: php
file: main.php
description: |
  Runs an OpenFeature provider doc fragment that assumes a registered
  provider and a bound `$client` and `$context` already exist — the
  "construct a context" and "evaluate a context" fragments. The scaffold
  registers a real LaunchDarkly provider, binds `$provider`, `$client`,
  and a default `$context`, then runs the fragment (which may reassign
  `$context`) and evaluates a flag, printing the success line. Requires
  LaunchDarkly credentials because the provider connects.
inputs:
  body:
    type: string
    description: The wrappee fragment, run with `$provider`, `$client`, and `$context` in scope.
validation:
  runtime: php
  entrypoint: main.php
  requirements: |
    launchdarkly/openfeature-server
---

```php
<?php
require __DIR__ . '/vendor/autoload.php';

use OpenFeature\OpenFeatureAPI;
use OpenFeature\implementation\flags\Attributes;
use OpenFeature\implementation\flags\EvaluationContext;

$provider = new LaunchDarkly\OpenFeature\Provider(getenv('LAUNCHDARKLY_SDK_KEY'), []);
$api = OpenFeatureAPI::getInstance();
$api->setProvider($provider);
$client = $api->getClient("hello-client", 1);
$context = new EvaluationContext("example-user-key");

{{ body }}

$flagValue = $client->getBooleanValue(getenv('LAUNCHDARKLY_FLAG_KEY'), false, $context);

echo "feature flag evaluates to true\n";
```
