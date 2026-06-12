---
id: haskell-server-sdk/sdk-docs/features/testdata/configure
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Test data source configuration for Haskell.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only
---

```haskell
import qualified LaunchDarkly.Server.Integrations.TestData as TestData

td <- TestData.newTestData
let config = LD.configSetDataSourceFactory (Just $ TestData.dataSourceFactory td) $ LD.makeConfig "YOUR_SDK_KEY"
client <- LD.makeClient config
```
