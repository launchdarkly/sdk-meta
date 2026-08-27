---
id: php-server-sdk/sdk-info/init
sdk: php-server-sdk
kind: init
lang: php
file: php-server-sdk/init.txt
description: |
  Client initialization snippet for php-server-sdk. A complete runnable
  program: the PHP SDK evaluates flags synchronously per request, so
  constructing LDClient is all the initialization it needs (there is no
  async connect to await). The wrappee owns the whole program, so the
  harness asserts the snippet's own success line via SNIPPET_SUCCESS_RE.
validation:
  runtime: php
  requirements: |
    launchdarkly/server-sdk
  placeholders:
    YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
  env:
    SNIPPET_SUCCESS_RE: SDK successfully initialized
---

```php
<?php

require 'vendor/autoload.php';

// This is your LaunchDarkly SDK key.
// Never hardcode your SDK key in production.
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY");

// The PHP SDK evaluates flags synchronously per request, so constructing
// the client is all the initialization it needs.
echo "SDK successfully initialized\n";
```
