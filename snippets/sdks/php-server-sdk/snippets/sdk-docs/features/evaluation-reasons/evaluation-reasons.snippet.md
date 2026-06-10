---
id: php-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons
sdk: php-server-sdk
kind: reference
lang: php
description: Flag evaluation reason example for PHP.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$detail = $client->variationDetail("example-flag-key", $myContext, false);

$value = $detail->getValue();
$index = $detail->getVariationIndex();
$reason = $detail->getReason();
```
