---
id: php-server-sdk/sdk-docs/features/testdata/configure
sdk: php-server-sdk
kind: reference
lang: php
description: Test data source configuration for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
require 'vendor/autoload.php';

$td = new LaunchDarkly\Integrations\TestData();
// You can set any initial flag states here with td.update

$client = new LaunchDarkly\LDClient($sdkKey, ['feature_requester' => $td]);
```
