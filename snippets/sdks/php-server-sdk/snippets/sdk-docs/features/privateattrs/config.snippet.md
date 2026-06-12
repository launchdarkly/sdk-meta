---
id: php-server-sdk/sdk-docs/features/privateattrs/config
sdk: php-server-sdk
kind: reference
lang: php
description: Private attribute configuration for PHP SDK v5.0.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
// All attributes marked private
$client = new LaunchDarkly\LDClient($sdkKey, ['all_attributes_private' => true]);

// Two attributes marked private
$client = new LaunchDarkly\LDClient($sdkKey, ['private_attribute_names' => ['name', 'email']]);
```
