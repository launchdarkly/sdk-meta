---
id: php-server-sdk/sdk-docs/features/migrations/migration-variation
sdk: php-server-sdk
kind: reference
lang: php
description: Migration stage evaluation (migrationVariation) for PHP SDK v6.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
$context = LaunchDarkly\LDContext::builder("example-context-key")->build();

$result = $client->migrationVariation('example-migration-flag-key', $context, Migrations\Stage::OFF);

/** @var Migrations\Stage */
$stage = $result['stage'];
/** @var Migrations\OpTracker */
$tracker = $result['tracker'];
```
