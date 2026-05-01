---
id: php-server-sdk/sdk-docs/migration-5-to-6-configuring-the-migration-php-sdk-v6
sdk: php-server-sdk
kind: reference
lang: php
description: "PHP SDK v6 in section \"Configuring the migration\""
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
