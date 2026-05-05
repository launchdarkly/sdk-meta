---
id: js-client-sdk/sdk-docs/requirements-polyfills-document-queryselectorall-html
sdk: js-client-sdk
kind: reference
lang: html
description: "HTML in section \"document.querySelectorAll()\""
# Bucket C: HTML <script src="…"> tag, not JS code; the
# js-syntax-only scaffold writes the body as TypeScript and tsdown
# rejects raw HTML. See _sdk-docs-port-notes.md.
---

```html
<!-- loading polyfill from CDN -->
<script src="https://unpkg.com/polyfill-queryselector@1.0.2/querySelector.js"></script>
```
