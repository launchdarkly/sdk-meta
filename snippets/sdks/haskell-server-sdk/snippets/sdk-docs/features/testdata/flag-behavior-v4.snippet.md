---
id: haskell-server-sdk/sdk-docs/features/testdata/flag-behavior-v4
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Configuring test data flag behavior for Haskell SDK v4.0.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only
---

```haskell
-- This flag is true for the context with kind "context-kind" and key "example-context-key", and false for everyone else
TestData.update td =<< ( TestData.flag td "flag-key-456def"
	<&> TestData.booleanFlag
	<&> TestData.variationForKey "context-kind" "example-context-key" (0 :: TestData.VariationIndex)
	<&> TestData.fallthroughVariation (1 :: TestData.VariationIndex)
)

-- This flag returns the string variation "green" for contexts with kind "context-kind"
-- that have the custom attribute "admin" with a value of true, and "red" for everyone else.
TestData.update td =<< ( TestData.flag td "flag-key-789ghi"
	<&> TestData.variations [toJSON "red", toJSON "green"]
	<&> TestData.ifMatchContext "context-kind" "admin" [Aeson.Bool True]
	<&> TestData.thenReturn (1 :: TestData.VariationIndex)
	<&> TestData.fallthroughVariation (0 :: TestData.VariationIndex)
)
```
