---
id: haskell-server-sdk/sdk-docs/features/testdata/flag-behavior-v3x
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Configuring test data flag behavior for Haskell SDK v3.x.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only-v3
---

```haskell
-- This flag is true for the user key "example-user-key" and false for everyone else
TestData.update td =<< ( TestData.flag td "flag-key-456def"
	<&> TestData.booleanFlag
	<&> TestData.variationForUser "example-user-key" (0 :: TestData.VariationIndex)
	<&> TestData.fallthroughVariation (1 :: TestData.VariationIndex)
)

-- This flag returns the string variation "green" for users who have the custom
-- attribute "admin" with a value of true, and "red" for everyone else.
TestData.update td =<< ( TestData.flag td "flag-key-456def"
	<&> TestData.variations [toJSON "red", toJSON "green"]
	<&> TestData.ifMatch "admin" [Aeson.Bool True]
	<&> TestData.thenReturn (1 :: TestData.VariationIndex)
	<&> TestData.fallthroughVariation (0 :: TestData.VariationIndex)
)
```
