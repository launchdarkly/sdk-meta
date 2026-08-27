---
id: roku-client-sdk/sdk-info/install
sdk: roku-client-sdk
kind: install
lang: shell
file: roku-client-sdk/install.txt
description: |
  Install snippet for roku-client-sdk. The SDK has no package manager;
  it ships as a package.zip asset on GitHub releases and the files are
  copied into the channel's source tree. Unvalidated: shell-install has
  no curl/unzip support, and there is nothing to resolve from a registry.
---

```shell
curl -L -o package.zip https://github.com/launchdarkly/roku-client-sdk/releases/latest/download/package.zip
unzip package.zip

# Copy the SDK files into your channel's source tree:
#   LaunchDarkly.brs      -> source/
#   LaunchDarklyTask.brs  -> components/
#   LaunchDarklyTask.xml  -> components/
```
