---
id: php-server-sdk/sdk-docs/features/hooks/define-hook
sdk: php-server-sdk
kind: reference
lang: php
description: Hook implementation and configuration for the PHP SDK.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
use LaunchDarkly\Hooks\EvaluationSeriesContext;
use LaunchDarkly\Hooks\Hook;
use LaunchDarkly\Hooks\Metadata;
use LaunchDarkly\Hooks\TrackSeriesContext;
use LaunchDarkly\EvaluationDetail;
use LaunchDarkly\LDClient;

class ExampleHook extends Hook
{
    public function getMetadata(): Metadata
    {
        return new Metadata('example-hook');
    }

    // Implement at least one of `beforeEvaluation`, `afterEvaluation`, or `afterTrack`

    // `beforeEvaluation` is called during the execution of a variation method
    // before the flag value has been determined

    // `afterEvaluation` is called during the execution of a variation method
    // after the flag value has been determined

    // `afterTrack` is called after a custom event has been enqueued by `track`
}

$client = new LDClient('YOUR_SDK_KEY', [
    'hooks' => [new ExampleHook()],
]);
```
