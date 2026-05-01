---
id: php-server-sdk/sdk-docs/migration-4-to-5-understanding-differences-between-users-and-contexts-5-0-syntax-multi-context
sdk: php-server-sdk
kind: reference
lang: php
description: "5.0 syntax, multi-context in section \"Understanding differences between users and contexts\""
---

```php
$deviceContext = LDContext::create("example-device-key", "device");
$orgContext = LDContext::create("example-organization-key", "org");
$multiContext = LDContext::createMulti($deviceContext, $orgContext);
```
