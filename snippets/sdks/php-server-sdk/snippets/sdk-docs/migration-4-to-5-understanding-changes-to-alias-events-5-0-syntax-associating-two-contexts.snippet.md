---
id: php-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-alias-events-5-0-syntax-associating-two-contexts
sdk: php-server-sdk
kind: reference
lang: php
description: "5.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```php
$deviceContext = LDContext::create("example-device-key", "device");
$orgContext = LDContext::create("example-organization-key", "org");
$multiContext = LDContext::createMulti($deviceContext, $orgContext);

$client->identify($multiContext);
```
