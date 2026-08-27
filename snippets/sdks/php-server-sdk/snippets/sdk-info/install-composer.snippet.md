---
id: php-server-sdk/sdk-info/install-composer
sdk: php-server-sdk
kind: install
lang: shell
file: php-server-sdk/install-composer.txt
description: |
  Install command for php-server-sdk (composer). guzzlehttp/guzzle is a
  suggested (not required) dependency of the SDK package, but the
  default FeatureRequester — which LDClient constructs — hard-requires
  it, so onboarding installs need both.
validation:
  runtime: shell-install
---

```shell
composer require launchdarkly/server-sdk guzzlehttp/guzzle
```
