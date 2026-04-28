---
id: android-client-sdk/getting-started/manifest
sdk: android-client-sdk
kind: manifest-fragment
lang: xml
description: AndroidManifest.xml fragment registering MainApplication.
ld-application:
  slot: manifest
---

Register the `MainApplication` class in the `AndroidManifest.xml`:

```xml
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
  xmlns:tools="http://schemas.android.com/tools">

  <application
    android:name=".MainApplication"
    ...
    ...
  </application>
</manifest>
```
