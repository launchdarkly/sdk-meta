---
id: dotnet-server-sdk/sdk-docs/features/securemode/compute-hash-manually
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Manual secure mode hash computation example in C#.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp
using System;
using System.Security.Cryptography;
using System.Text;

var encoding = new UTF8Encoding();
var keyBytes = encoding.GetBytes("YOUR_SDK_KEY");
var hmacSha256 = new HMACSHA256(keyBytes);
var hashBytes = hmacSha256.ComputeHash(encoding.GetBytes("example-context-key"));
var hashString = BitConverter.ToString(hashBytes).Replace("-", "").ToLower();
```
