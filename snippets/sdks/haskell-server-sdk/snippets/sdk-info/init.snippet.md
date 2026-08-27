---
id: haskell-server-sdk/sdk-info/init
sdk: haskell-server-sdk
kind: init
lang: haskell
file: haskell-server-sdk/init.txt
description: |
  Client initialization snippet for haskell-server-sdk. A complete
  runnable Main.hs (modeled on launchdarkly/hello-haskell-server): make
  the client, poll `getStatus` until it leaves `Uninitialized`, and
  report success or failure. Validated end-to-end by the haskell-server
  harness; because the body owns `main`, the harness asserts the
  snippet's own success line via `SNIPPET_SUCCESS_RE` instead of
  wrapping it in a trailer-emitting scaffold.
validation:
  runtime: haskell-server
  placeholders:
    YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
  env:
    SNIPPET_SUCCESS_RE: SDK successfully initialized
---

```haskell
{-# LANGUAGE OverloadedStrings #-}

module Main where

import Control.Concurrent (threadDelay)
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
    -- This is your LaunchDarkly SDK key.
    -- Never hardcode your SDK key in production.
    client <- LD.makeClient $ LD.makeConfig "YOUR_SDK_KEY"

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
