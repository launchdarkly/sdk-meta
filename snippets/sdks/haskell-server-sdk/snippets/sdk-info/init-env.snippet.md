---
id: haskell-server-sdk/sdk-info/init-env
sdk: haskell-server-sdk
kind: init
lang: haskell
file: haskell-server-sdk/init-env.txt
description: |
  Client initialization snippet for haskell-server-sdk reading the SDK
  key from an environment variable. Uses base's `Data.String.fromString`
  (rather than `Data.Text.pack`) to build the config's Text key, so the
  snippet compiles in projects that only list launchdarkly-server-sdk in
  their direct dependencies. No `placeholders:` needed — the harness
  already exports LAUNCHDARKLY_SDK_KEY into the validate run.
validation:
  runtime: haskell-server
  env:
    SNIPPET_SUCCESS_RE: SDK successfully initialized
---

```haskell
module Main where

import Control.Concurrent (threadDelay)
import Data.String (fromString)
import System.Environment (getEnv)
import System.Exit (exitFailure)
import System.Timeout (timeout)
import qualified LaunchDarkly.Server as LD

waitForClient :: LD.Client -> IO Bool
waitForClient client = do
    status <- LD.getStatus client
    case status of
        LD.Uninitialized -> threadDelay 1000 >> waitForClient client
        LD.Initialized -> pure True
        _ -> pure False

main :: IO ()
main = do
    -- Reads your LaunchDarkly SDK key from the LAUNCHDARKLY_SDK_KEY
    -- environment variable, so one build can run in every environment.
    sdkKey <- getEnv "LAUNCHDARKLY_SDK_KEY"
    client <- LD.makeClient $ LD.makeConfig (fromString sdkKey)

    -- Wait up to five seconds for the client to connect to LaunchDarkly.
    initialized <- timeout (5 * 1000 * 1000) (waitForClient client)
    case initialized of
        Just True -> do
            -- For onboarding purposes only we flush events as soon as
            -- possible so we quickly detect your connection.
            -- You don't have to do this in practice because events are automatically flushed.
            LD.flushEvents client
            putStrLn "SDK successfully initialized"
        _ -> do
            putStrLn "SDK failed to initialize"
            exitFailure
```
