---
id: php-server-sdk/sdk-docs/features/migrations/stage-switch
sdk: php-server-sdk
kind: reference
lang: php
description: Per-stage migration structure for PHP SDK v6.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
// define the combination of reads and writes from the new and old systems
// that should occur at each migration stage

switch ($stage) {
  case Migrations\Stage::OFF:
  case Migrations\Stage::DUALWRITE:
  case Migrations\Stage::SHADOW:
  case Migrations\Stage::LIVE:
  case Migrations\Stage::RAMPDOWN:
  case Migrations\Stage::COMPLETE:
  default:
    // throw an error
}
```
