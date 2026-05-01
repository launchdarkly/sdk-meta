---
id: react-native-client-sdk/sdk-docs/initialize-the-client-and-identify-a-context-react-native-sdk-v10-x
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: "React Native SDK v10.x in section \"Initialize the client and identify a context\""
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
