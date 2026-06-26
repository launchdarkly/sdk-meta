# Port notes: /sdk/features/data-saving-mode

Source: `ld-docs-private` `fern/topics/sdk/features/data-saving-mode.mdx`.
29 code blocks extracted into `sdk-docs/features/datasaving/` snippets
across 10 SDKs. All 29 are bound to validators — this page has no
Objective-C section, so there are zero non-binds.

## Content corrections

None. Every body is verbatim from the MDX. The page documents the EAP
data system surface; every API call was verified against the local SDK
repos before binding:

- .NET: `Components.DataSystem()`, `DataSystemComponents.Polling()/.Streaming()`,
  `ServiceEndpointsOverride`, `FileData.DataSource().FilePaths`,
  `ConfigurationBuilder.DataSystem` (dotnet-core, shipped in 8.11.0).
- Go: `Config.DataSystem`, `ldcomponents.DataSystem().Default()/.Custom()`,
  `PollingDataSourceV2`/`StreamingDataSourceV2` with `BaseURI(string)` /
  `AsInitializer()`, `ldfiledatav2.DataSource().FilePaths` (go-server-sdk,
  shipped in 7.11.0).
- Java: `Components.dataSystem().defaultMode()/.custom()`,
  `DataSystemComponents.pollingInitializer()/streamingSynchronizer()/pollingSynchronizer()`,
  `serviceEndpointsOverride(ServiceEndpointsBuilder)`,
  `FileData.initializer().filePaths` (java-core, shipped in 7.11.0).
- Node.js: `dataSystem.dataSource.dataSourceOptionsType: 'standard' | 'custom'`,
  tagged initializer/synchronizer unions `{type: 'file'|'polling'|'streaming'}`
  with `paths` (js-core server-node, shipped in 9.10).
- Python: `ldclient.datasystem` `default()/custom()/polling_ds_builder()/
  streaming_ds_builder()/file_ds_builder(paths=[...])`; `initializers`
  takes a list while `synchronizers` is varargs — the page's asymmetric
  calls match the real signatures (python-server-sdk, shipped in 9.13.0).
- Ruby: `LaunchDarkly::DataSystem.default/custom/polling_ds_builder/
  streaming_ds_builder`, `Integrations::FileData.data_source_v2(paths:)`,
  both `initializers`/`synchronizers` take arrays (ruby-server-sdk,
  shipped in 8.12.0).
- Android: `Components.dataSystem()`, `DataSystemBuilder.automaticModeSwitching/
  foregroundConnectionMode`, `AutomaticModeSwitchingConfig.disabled()`,
  `DataSystemComponents.automaticModeSwitching().lifecycle(...).network(...).build()`,
  `ConnectionMode.STREAMING` (android-client-sdk, shipped in 5.13.0 EAP).
- JS (browser): `createClient(clientSideId, context, options)` with
  `dataSystem?: BrowserDataSystemOptions`; `automaticModeSwitching` is
  `false | ManualModeSwitching` where ManualModeSwitching is
  `{type: 'manual', initialConnectionMode}` (js-core browser, 4.9.0).
- React Native: options `dataSystem?: RNDataSystemOptions` whose
  `automaticModeSwitching` accepts `false` (js-core react-native, 10.19.0).
- React Web: `createLDReactProvider(clientSideID, context, {ldOptions})`
  with `ldOptions.dataSystem` (js-core react). Note: the local js-core
  checkout shows the data-saving release as react-sdk 4.1.2 while the
  page prose says "version 4.2 or later"; prose lives in the MDX, not in
  any snippet body, so nothing to correct here.

## Validation routing added in this port

- `validators/languages/android-client/Dockerfile`:
  `LD_ANDROID_SDK_VERSION` 5.11.1 --> 5.13.1. The data system surface
  first appears in the published 5.13.0 aar; the kotlin/java
  syntax-only scaffolds compile fragments against the real aar, so the
  six Android fragments need the newer pin. Existing android-bound
  snippets were re-validated against the bumped image.

No new scaffolds; all 29 snippets bind to pre-existing syntax-only
scaffolds (csharp, go, java/jvm, node, python, ruby, android kotlin +
java, js-client, react-native, react).

## Known non-binds

None.
