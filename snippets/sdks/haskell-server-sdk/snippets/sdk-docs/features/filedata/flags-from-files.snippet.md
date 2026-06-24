---
id: haskell-server-sdk/sdk-docs/features/filedata/flags-from-files
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: File data source configuration example for Haskell.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only

---

```haskell
let config = LD.configSetDataSourceFactory (Just $ FileData.dataSourceFactory ["./testData/flags.json"]) $ LD.makeConfig "YOUR_SDK_KEY"
client <- LD.makeClient config
```
