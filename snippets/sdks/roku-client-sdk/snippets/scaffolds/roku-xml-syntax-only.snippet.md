---
id: roku-client-sdk/scaffolds/roku-xml-syntax-only
sdk: roku-client-sdk
kind: scaffold
lang: xml
file: component.xml
description: |
  Well-formedness validator for Roku SceneGraph XML component doc
  fragments. Routes through the `xml` Docker validator, which runs
  `xmllint --noout` over the staged file. Parse-only: no SceneGraph
  schema validation, no Roku device, no LD env. Doc fragments are
  whole component files, so the scaffold emits the body verbatim.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, parsed as an XML document.
validation:
  runtime: xml
  entrypoint: component.xml
---

```xml
{{ body }}
```
