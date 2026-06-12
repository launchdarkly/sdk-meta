---
id: php-server-sdk/sdk-docs/features/testdata/flag-behavior-v5
sdk: php-server-sdk
kind: reference
lang: php
description: Configuring test data flag behavior for PHP SDK v5.0.
validation:
  scaffold: php-server-sdk/scaffolds/php-syntax-only
---

```php
// This flag is true for the context with the key "example-context-key" and kind of "organization",
// and false for everyone else
$td->update(
    $td->flag("example-flag-key")
       ->variationForKey("organization", "example-context-key", true)
       ->fallthroughVariation(false)
);

// This flag returns the string variation "green" for contexts that have the custom
// attribute "admin" with a value of true, and "red" for everyone else.
$td->update(
    $td->flag("example-flag-key")
       ->variations("red", "green")
       ->fallthroughVariation(0)
       ->ifMatch("admin", true)
       ->thenReturn(1)
);
```
