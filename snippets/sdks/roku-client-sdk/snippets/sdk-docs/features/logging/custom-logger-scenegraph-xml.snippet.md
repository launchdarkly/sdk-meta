---
id: roku-client-sdk/sdk-docs/features/logging/custom-logger-scenegraph-xml
sdk: roku-client-sdk
kind: reference
lang: xml
description: SceneGraph API custom logger component definition for Roku (CustomLogger.xml).
validation:
  scaffold: roku-client-sdk/scaffolds/roku-xml-syntax-only

---

```xml
<!-- /components/CustomLogger.xml -->

<component name="CustomLogger" extends="Task">
    <interface>
        <field id="log" type="assocarray" alwaysNotify="true"/>
    </interface>

    <script type="text/brightscript" uri="pkg:/components/CustomLogger.brs"/>
</component>
```
