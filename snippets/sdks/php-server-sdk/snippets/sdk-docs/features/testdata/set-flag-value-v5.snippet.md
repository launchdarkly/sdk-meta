---
id: php-server-sdk/sdk-docs/features/testdata/set-flag-value-v5
sdk: php-server-sdk
kind: reference
lang: php
description: Setting a test data flag to a specific value for PHP SDK v5.0.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
$td->update($td->flag("example-flag-key")->variationForAll(false));
```
