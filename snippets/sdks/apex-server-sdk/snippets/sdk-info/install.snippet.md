---
id: apex-server-sdk/sdk-info/install
sdk: apex-server-sdk
kind: install
lang: shell
file: apex-server-sdk/install.txt
description: |
  Install snippet for apex-server-sdk (Salesforce CLI deploy). The SDK
  has no package manager; it deploys from a clone of the SDK repo.
  Unvalidated: shell-install has no sf/git support and there is no
  Salesforce org in CI.
---

```shell
git clone https://github.com/launchdarkly/apex-server-sdk.git
cd apex-server-sdk
sf project deploy start --target-org 'YOUR TARGET ORG' --source-dir 'force-app'
```
