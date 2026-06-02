---
id: dotnet-client-sdk/sdk-docs/install-the-sdk-c
sdk: dotnet-client-sdk
kind: reference
lang: csharp
description: "C# in section \"Install the SDK\""
# TODO(snippet-bug): body is `Install-Package LaunchDarkly.ClientSdk`
# — Visual Studio's NuGet Package Manager Console PowerShell command,
# mistagged `csharp` in the source MDX. No C# parser will accept it.
# Fix in the snippet-bugs PR: either re-tag as `powershell` (and
# build a corresponding parse path) or split into separate
# `dotnet add package …` / `Install-Package …` snippets, each
# correctly tagged.
---

```csharp
Install-Package LaunchDarkly.ClientSdk
```
