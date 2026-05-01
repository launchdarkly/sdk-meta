---
id: php-server-sdk/scaffolds/php-syntax-only
sdk: php-server-sdk
kind: scaffold
lang: php
file: main.php
description: |
  Parse-only validator for PHP server SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: php
  entrypoint: main.php
---

```php
<?php
function _wrappee() {
{{ body }}
}
echo "feature flag evaluates to true\n";
```
