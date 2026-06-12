---
id: android-client-sdk/sdk-docs/features/monitoring/status-listener-java
sdk: android-client-sdk
kind: reference
lang: java
description: Status listener registration in an Activity for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
class MainActivity extends Activity {
    private LDClient client;
    private LDStatusListener ldStatusListener;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        ldStatusListener = new LDStatusListener() {
            @Override
            public void onConnectionModeChanged(ConnectionInformation connectionInfo) {
                // handle new connection info
            }

            @Override
            public void onInternalFailure(LDFailure ldFailure) {
                // handle failure
            }
        };

        try {
            client = LDClient.get();
        } catch (LaunchDarklyException e) {
            // client was not initialized before calling get()
            return;
        }
        client.registerStatusListener(ldStatusListener);
    }

    @Override
    protected void onDestroy() {
        super.onDestroy();
        client.unregisterStatusListener(ldStatusListener);
    }
}
```
