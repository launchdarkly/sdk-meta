---
id: php-server-sdk/getting-started/install
sdk: php-server-sdk
kind: install
lang: shell
description: Install the LaunchDarkly SDK and Guzzle via composer.
inputs:
  version:
    type: string
    description: Optional pinned SDK version; when empty the pin is omitted.
    runtime-default: ""
ld-application:
  slot: install
---

Next, install the LaunchDarkly SDK and Guzzle dependency:

```shell
php composer.phar require launchdarkly/server-sdk{{ if version }}:{{ version }}{{ end }} guzzlehttp/guzzle
```
