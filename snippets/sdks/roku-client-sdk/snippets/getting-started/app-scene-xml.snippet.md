---
id: roku-client-sdk/getting-started/app-scene-xml
sdk: roku-client-sdk
kind: manifest
lang: xml
file: components/AppScene.xml
description: SceneGraph component definition with the LaunchDarklyTask node.
ld-application:
  slot: app-scene-xml
---

In `components/AppScene.xml` create a basic scene by adding the following code:

```xml
<?xml version="1.0" encoding="utf-8" ?>
<component name="AppScene" extends="Scene">
    <children>
        <LaunchDarklyTask id="launchDarkly"/>

        <Label id="evaluation"
            text="waiting on payload to initialize"
            width="1280"
            height="720"
            wrap="true"
            horizAlign="center"
            vertAlign="center"
        />

        <Label id="status"
            text="waiting for sdk status report"
            width="1280"
            height="720"
            wrap="true"
            horizAlign="center"
            vertAlign="bottom"
        />

    </children>

    <script type="text/brightscript" uri="pkg:/components/AppScene.brs"/>

    <!-- Include the LaunchDarkly SDK. -->
    <script type="text/brightscript" uri="pkg:/source/LaunchDarkly.brs"/>
</component>
```
