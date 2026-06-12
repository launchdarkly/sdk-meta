---
id: php-server-sdk/sdk-docs/features/storing-data/dynamodb/dynamodb
sdk: php-server-sdk
kind: reference
lang: php
description: DynamoDB feature requester configuration example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$fr = LaunchDarkly\Integrations\DynamoDb::featureRequester([
    'dynamodb_table' => 'my-table'
]);
$client = new LaunchDarkly\LDClient("YOUR_SDK_KEY", [
    'feature_requester' => $fr
]);
```
