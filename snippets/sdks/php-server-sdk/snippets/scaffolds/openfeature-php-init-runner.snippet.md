---
id: php-server-sdk/scaffolds/openfeature-php-init-runner
sdk: php-server-sdk
kind: scaffold
lang: php
file: main.php
description: |
  Runs an OpenFeature "initialize the provider" fragment end-to-end
  against a real LaunchDarkly environment. The fragment is expected to
  register a LaunchDarkly provider with OpenFeature and bind a `$client`;
  the scaffold supplies the `use` imports the fragment assumes, then
  uses `$client` to evaluate a flag and print the success line. The
  fragment's `YOUR_SDK_KEY` literal is replaced with the real key via
  the snippet's `validation.placeholders` before this runs.
inputs:
  body:
    type: string
    description: The wrappee init fragment; registers the provider and binds `$client`.
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
use OpenFeature\implementation\flags\EvaluationContext;

{{ body }}

$context = new EvaluationContext("example-user-key");
$flagValue = $client->getBooleanValue(getenv('LAUNCHDARKLY_FLAG_KEY'), false, $context);

echo "feature flag evaluates to true\n";
```
