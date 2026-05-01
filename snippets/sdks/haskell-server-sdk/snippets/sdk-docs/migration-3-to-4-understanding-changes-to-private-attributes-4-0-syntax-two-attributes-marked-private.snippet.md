---
id: haskell-server-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-private-attributes-4-0-syntax-two-attributes-marked-private
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: "4.0 syntax, two attributes marked private in section \"Understanding changes to private attributes\""
---

```haskell
import qualified Data.Set as S
import qualified LaunchDarkly.Server.Reference as R

makeConfig sdkKey
  & configSetAllAttributesPrivate True
  & configSetPrivateAttributeNames (S.fromList $ map R.makeLiteral ["name", "email"])
config = LaunchDarkly::Config.new({private_attributes: ["name", "email"]})
```
