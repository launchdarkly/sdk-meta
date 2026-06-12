---
id: php-server-sdk/sdk-docs/features/migrations/read-write
sdk: php-server-sdk
kind: reference
lang: php
description: Migration read and write example for PHP SDK v6.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
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
