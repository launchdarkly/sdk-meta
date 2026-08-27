---
id: php-server-sdk/sdk-info/init-env
sdk: php-server-sdk
kind: init
lang: php
file: php-server-sdk/init-env.txt
description: Client initialization snippet for php-server-sdk reading the SDK key from an environment variable.
validation:
  runtime: php
  requirements: |
    launchdarkly/server-sdk
  env:
    SNIPPET_SUCCESS_RE: SDK successfully initialized
---

```php
<?php

require 'vendor/autoload.php';

// Reads your LaunchDarkly SDK key from the LAUNCHDARKLY_SDK_KEY
// environment variable, so one build can run in every environment.
$client = new LaunchDarkly\LDClient(getenv("LAUNCHDARKLY_SDK_KEY"));

// The PHP SDK evaluates flags synchronously per request, so constructing
// the client is all the initialization it needs.
echo "SDK successfully initialized\n";
```
