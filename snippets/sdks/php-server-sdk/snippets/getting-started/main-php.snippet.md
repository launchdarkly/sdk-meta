---
id: php-server-sdk/getting-started/main-php
sdk: php-server-sdk
kind: hello-world
lang: php
file: main.php
description: Hello-world program that initializes the PHP server SDK and watches a feature flag.
inputs:
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source. Validation reads LAUNCHDARKLY_FLAG_KEY at runtime.
ld-application:
  slot: main-php
validation:
  runtime: php
  requirements: |
    launchdarkly/server-sdk
    guzzlehttp/guzzle
---

Create a file called `main.php` and add the following code:

```php
<?php

require 'vendor/autoload.php';

function showEvaluationResult(string $key, bool $value) {
  echo PHP_EOL;
  echo sprintf("*** %s: The %s feature flag evaluates to %s", date("h:i:s"), $key, $value ? 'true' : 'false');
  echo PHP_EOL;
}

function showBanner() {
  echo PHP_EOL;
  echo "        ██       " . PHP_EOL;
  echo "          ██     " . PHP_EOL;
  echo "      ████████   " . PHP_EOL;
  echo "         ███████ " . PHP_EOL;
  echo "██ LAUNCHDARKLY █" . PHP_EOL;
  echo "         ███████ " . PHP_EOL;
  echo "      ████████   " . PHP_EOL;
  echo "          ██     " . PHP_EOL;
  echo "        ██       " . PHP_EOL;
  echo PHP_EOL;
}

// Set $sdkKey to your LaunchDarkly SDK key.
$sdkKey = getenv("LAUNCHDARKLY_SDK_KEY") ?? "";

// Set $featureFlagKey to the feature flag key you want to evaluate.
$featureFlagKey = "{{ featureKey }}";


if (!$sdkKey) {
echo "*** Please set the environment variable LAUNCHDARKLY_SDK_KEY to your LaunchDarkly SDK key first" . PHP_EOL . PHP_EOL;
exit(1);
} else if (!$featureFlagKey) {
echo "*** Please set the environment variable LAUNCHDARKLY_FLAG_KEY to a boolean flag first" . PHP_EOL . PHP_EOL;
exit(1);
}

$client = new LaunchDarkly\LDClient($sdkKey);

// Set up the evaluation context. This context should appear on your LaunchDarkly contexts dashboard soon after you run the demo.
$context = LaunchDarkly\LDContext::builder("example-user-key")
->kind("user")
->name("Sandy")
->build();


$showBanner = true;
$lastValue = null;
do {
  $flagValue = $client->variation($featureFlagKey, $context, false);

  if ($flagValue !== $lastValue) {
      showEvaluationResult($featureFlagKey, $flagValue);
  }

  if ($showBanner && $flagValue) {
      showBanner();
      $showBanner = false;
  }

  $lastValue = $flagValue;
  sleep(1);
} while(true);
```
