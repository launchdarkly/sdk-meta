---
id: haskell-server-sdk/sdk-docs/features/privateattrs/config-v4
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Private attribute configuration for Haskell SDK v4.0.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only-toplevel

---

```haskell
import qualified Data.Set as S
import qualified LaunchDarkly.Server.Reference as R

-- All attributes marked private
configAllPrivate = makeConfig "YOUR_SDK_KEY" & configSetAllAttributesPrivate True

-- Two attributes marked private
configSomePrivate = makeConfig sdkKey
  & configSetPrivateAttributeNames (S.fromList $ map R.makeLiteral ["name", "email"])
```
