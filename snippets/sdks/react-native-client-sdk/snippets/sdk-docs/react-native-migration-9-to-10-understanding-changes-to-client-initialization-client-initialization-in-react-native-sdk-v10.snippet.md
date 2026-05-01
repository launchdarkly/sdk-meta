---
id: react-native-client-sdk/sdk-docs/react-native-migration-9-to-10-understanding-changes-to-client-initialization-client-initialization-in-react-native-sdk-v10
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: "Client initialization in React Native SDK v10 in section \"Understanding changes to client initialization\""
---

```ts
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
