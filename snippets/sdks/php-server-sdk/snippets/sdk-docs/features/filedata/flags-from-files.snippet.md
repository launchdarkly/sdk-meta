---
id: php-server-sdk/sdk-docs/features/filedata/flags-from-files
sdk: php-server-sdk
kind: reference
lang: php
description: File data source configuration example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
// Automatic reloading is not supported in PHP, because normally in PHP
// the entire in-memory application state is recreated for each request.

$fr = LaunchDarkly\Integrations\Files::featureRequester([
    'file1.json',
    'file2.json'
]);
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", [
    'feature_requester' => $fr,
    'send_events' => false
]);
```
