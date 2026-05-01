---
id: php-server-sdk/sdk-docs/migration-5-to-6-reading-and-writing-during-the-migration-php-sdk-v6
sdk: php-server-sdk
kind: reference
lang: php
description: "PHP SDK v6 in section \"Reading and writing during the migration\""
---

```php

$context = LaunchDarkly\LDContext::builder("example-context-key")->build();

// this is the migration stage to use if the flag's migration stage
// is not available from LaunchDarkly
$defaultStage = Migrations\Stage::OFF;

$result = $builder->build();
if (!$result->isSuccessful()) {
    throw new \Exception($result->error);
}

$migrator = $result->value;

// if you need to pass additional information from the call site
// to your read/write methods, use a mixed type payload
$payload = ['index' => 'useful information'];

// when you need to perform a read in your application
$migrator->read('example-migration-flag-key', $context, $defaultStage, $payload);

// when you need to perform a write in your application
$migrator->write('example-migration-flag-key', $context, $defaultStage, $payload);
```
