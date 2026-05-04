---
id: dotnet-server-sdk/scaffolds/init-runner-csproj
sdk: dotnet-server-sdk
kind: scaffold
lang: xml
file: HelloDotNet.csproj
description: |
  Companion `.csproj` for the dotnet-server init scaffold. Uses the
  ASP.NET Core SDK so `WebApplication.CreateBuilder` resolves; the
  harness's run.sh skips synthesizing its default console-only csproj
  when one is already staged.
---

```xml
<Project Sdk="Microsoft.NET.Sdk.Web">
  <PropertyGroup>
    <OutputType>Exe</OutputType>
    <TargetFramework>net8.0</TargetFramework>
    <Nullable>disable</Nullable>
    <ImplicitUsings>enable</ImplicitUsings>
    <RootNamespace>HelloDotNet</RootNamespace>
    <AssemblyName>HelloDotNet</AssemblyName>
  </PropertyGroup>
</Project>
```
