---
id: flutter-client-sdk/getting-started/main-dart
sdk: flutter-client-sdk
kind: hello-world
lang: dart
file: lib/main.dart
description: Hello-world Flutter app that initializes the LaunchDarkly SDK and renders the flag value.
inputs:
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source.
ld-application:
  slot: main-dart
# Validator pending — flutter test/run on a Linux runner with
# flutter-action; deferred.
---

Open the file `lib/main.dart` and replace with the following code:

```dart
import 'package:flutter/material.dart';
import 'package:launchdarkly_flutter_client_sdk/launchdarkly_flutter_client_sdk.dart';
import 'package:provider/provider.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    // The LDClient doesn't need to change throughout the lifetime of the
    // application, so we wrap the application in a provider with the client.
    return Provider<LDClient>(
        create: (_) => LDClient(
            LDConfig(
              // The credentials come from the environment, you can set them
              // using --dart-define.
              // Examples:
              // flutter run --dart-define LAUNCHDARKLY_CLIENT_SIDE_ID=<my-client-side-id> -d Chrome
              // flutter run --dart-define LAUNCHDARKLY_MOBILE_KEY=<my-mobile-key> -d ios
              //
              // Alternatively `CredentialSource.fromEnvironment()` can be replaced with your mobile key.
              CredentialSource.fromEnvironment(),
              AutoEnvAttributes.enabled,
            ),
            // Here we are using a default user with key of 'example-user-key'.
            LDContextBuilder().kind('user', 'example-user-key')
              .setString('name', 'Sandy').build()),
        dispose: (_, client) => client.close(),
        // We use a future provider to wait for the client to either start,
        // or for a timeout to elapse.
        child: MaterialApp(
          title: 'LaunchDarkly Hello App',
          theme: ThemeData(
            useMaterial3: true,
          ),
          home: const MyHomePage(),
        ));
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

/// Example provider which listens for flag changes and maps them to bool
/// values. It would also be possible to map to some application specific model
/// types. When mapping be sure all values are accessed through the client
/// `variation` methods. This ensures that the SDK generates the expected
/// events.
class FlagProviderBool extends StreamProvider<bool> {
  FlagProviderBool(
      {super.key,
      required LDClient client,
      required String flagKey,
      required bool defaultValue,
      required Widget child})
      : super(
            create: (context) => client.flagChanges
                .where((element) => element.keys.contains(flagKey))
                .map((event) => client.boolVariation(flagKey, defaultValue)),
            // Here we get the initial value of the flag. If the SDK is not
            // initialized, then the default value will be returned.
            initialData: client.boolVariation(flagKey, defaultValue),
            child: child);
}

class _MyHomePageState extends State<MyHomePage> {
  static const String flagKey = '{{ featureKey }}';

  @override
  Widget build(BuildContext context) {
    // The FutureBuilder here is used to gate the presentation content
    // based on the LaunchDarkly SDK having started. While it has not started,
    // a loading indicator will be shown. After it has started, or encountered
    // a timeout, then it will render the content.
    return FutureBuilder(
        future: Provider.of<LDClient>(context, listen: false)
            .start()
            // In this case we do not have special handling for a failed
            // initialization or timeout.
            .timeout(const Duration(seconds: 5), onTimeout: () => true)
            .then((value) => true),
                builder: (context, loaded) => loaded.data ?? false ?
                    FlagProviderBool(
                        // The client will not be changing, so we don't need to
                        // listen for client changes.
                        client: Provider.of<LDClient>(context, listen: false),
                        flagKey: flagKey,
                        defaultValue: false,
                        child: Consumer<bool>(
                          builder: (context, flagValue, _) => Scaffold(
                              backgroundColor: flagValue ? const Color(0xFF00844B) : const Color(0xFF373841),
                              body:
                                  Center(
                                    child: Text(
                                      'The $flagKey feature flag evaluates to $flagValue',
                                      style: const TextStyle(color: Colors.white, fontSize: 16)
                                    )
                                  )),
                        )) : const CircularProgressIndicator());
  }
}
```
