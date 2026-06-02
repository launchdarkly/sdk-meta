---
id: react-native-client-sdk/sdk-docs/initialize-the-client-and-identify-a-context-react-native-sdk-v10-x
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: "React Native SDK v10.x in section \"Initialize the client and identify a context\""
# TODO(validate): react-native validator harness expects a 2-file shape
# (App.tsx + src/welcome.tsx) baked into the prebuilt jest project,
# whereas the react-native-syntax-only scaffold writes a single App.js.
# The harness needs a separate parse-only mode (or the scaffold needs
# to output the App.tsx + src/welcome.tsx pair) before sdk-docs
# fragments can validate. See _sdk-docs-port-notes.md.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only
---

```ts
const options = {
  // optional observability plugin, requires React Native SDK v10.10+
  plugins: [ new Observability() ],
  // other options
}

const client = new ReactNativeLDClient('example-mobile-key', AutoEnvAttributes.Enabled, options);

const context = { kind: 'user', key: 'example-user-key' }

const App = () => {
  // call identify on App mount or later in some other component
  useEffect(() => {
    client.identify(context).catch((e: any) => console.log(e));
  }, []);

  return (
    <LDProvider client={client}>
      {/* your application code here */}
      <YourComponent />
    </LDProvider>
  );
};

export default App;
```
