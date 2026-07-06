---
id: php-server-sdk/scaffolds/openfeature-php-toplevel
sdk: php-server-sdk
kind: scaffold
lang: php
file: main.php
description: |
  Resolves an OpenFeature provider doc fragment that is a set of `use`
  imports. A bare `use` statement in PHP is only an alias and does not
  trigger autoloading, so the scaffold follows the imports with
  `class_exists`/`interface_exists` checks against the autoloader to
  confirm each imported name actually resolves to a class shipped by
  the installed OpenFeature packages. A typo in a namespace fails the
  check.
inputs:
  body:
    type: string
    description: The wrappee's `use` statements, staged verbatim at file scope.
validation:
  runtime: php
  entrypoint: main.php
  requirements: |
    launchdarkly/openfeature-server
---

```php
<?php
require __DIR__ . '/vendor/autoload.php';

{{ body }}

foreach ([
    'OpenFeature\OpenFeatureAPI',
    'OpenFeature\implementation\flags\Attributes',
    'OpenFeature\implementation\flags\EvaluationContext',
] as $fqcn) {
    if (!class_exists($fqcn) && !interface_exists($fqcn)) {
        fwrite(STDERR, "unresolved import: $fqcn\n");
        exit(1);
    }
}

echo "feature flag evaluates to true\n";
```
