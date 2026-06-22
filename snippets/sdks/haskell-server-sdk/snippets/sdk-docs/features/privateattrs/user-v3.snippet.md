---
id: haskell-server-sdk/sdk-docs/features/privateattrs/user-v3
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Marking user attributes private for Haskell SDK v3.x.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only-toplevel-v3

---

```haskell
import Data.Set (fromList)

user' = userSetPrivateAttributeNames (fromList ["name", "email"]) user
```
