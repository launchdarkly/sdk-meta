---
id: js-client-sdk/getting-started/index-html
sdk: js-client-sdk
kind: hello-world
lang: html
file: index.html
description: Static host page that loads the bundled SDK app.
ld-application:
  slot: index-html
---

Create `index.html`:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge;chrome=1">
    <title>LaunchDarkly tutorial</title>
    <script src="./dist/app.js" defer></script>
  </head>
  <body></body>
</html>
```
