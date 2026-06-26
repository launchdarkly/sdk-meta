---
id: haskell-server-sdk/sdk-docs/features/storing-data/redis/redis
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Redis feature store configuration example for Haskell.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only-module

---

```haskell
import qualified Database.Redis as               R
import           LaunchDarkly.Server
import           LaunchDarkly.Server.Store.Redis

main = do
    con <- R.checkedConnect R.defaultConnectInfo { R.connectHost = "my-redis", R.connectPort = R.PortNumber 6379 }
    backend <- makeRedisStore $ redisConfigSetNamespace "my-key-prefix" $ makeRedisStoreConfig con

    let config = configSetStoreBackend (Just backend) $ makeConfig "YOUR_SDK_KEY"

    client <- makeClient config
```
