---
id: php-server-sdk/sdk-docs/features/contextconfig/multi-context
sdk: php-server-sdk
kind: reference
lang: php
description: Multi-context example for PHP SDK v5.0.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only

---

```php
$deviceContext = LDContext::create("example-device-key", "device");
$orgContext = LDContext::create("example-organization-key", "org");
$multiContext = LDContext::createMulti($deviceContext, $orgContext);
```
