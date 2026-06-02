---
id: php-server-sdk/sdk-docs/using-guzzle-php
sdk: php-server-sdk
kind: reference
lang: php
description: "PHP in section \"Using Guzzle\""
# Bucket C: body is shell composer commands mistagged as `lang: php`
# in the source MDX. The php-syntax-only scaffold parses PHP, so this
# binding can't validate this body until either (a) the source MDX
# fence is retagged as shell (and the body adjusted to drop the
# leading `php` prefix so shell-install can sniff `composer`), or
# (b) a php-or-shell-composer scaffold is added. Tracked for the
# follow-up snippet-bugs PR — see _sdk-docs-port-notes.md.
---

```php
php composer.phar require "guzzlehttp/guzzle:^6.3.0"
php composer.phar require "kevinrob/guzzle-cache-middleware:^1.4.0"
```
