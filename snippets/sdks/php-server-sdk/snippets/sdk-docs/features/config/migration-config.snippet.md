---
id: php-server-sdk/sdk-docs/features/config/migration-config
sdk: php-server-sdk
kind: reference
lang: php
description: Migration configuration example for the PHP SDK v6 — read/write methods, execution order, latency/error tracking.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php

use LaunchDarkly\Migrations;
use LaunchDarkly\Types;

$builder = new Migrations\MigratorBuilder($client);

$builder->read(
    fn (?string $payload) => Types\Result::success("old read"),
    fn (?string $payload) => Types\Result::success("new read"),
    fn(string $old, string $new) => $old == $new,
);

$builder->write(
    fn (?string $payload) => Types\Result::success("old write"),
    fn (?string $payload) => Types\Result::success("new write")
);
$builder->readExecutionOrder(Migrations\ExecutionOrder::SERIAL);
// could also use ExecutionOrder::RANDOM

$builder->trackLatency(true); // defaults to true
$builder->trackErrors(true);  // defaults to true

$result = $builder->build();

```
