---
id: roku-client-sdk/sdk-docs/evaluate-a-flag-brightscript-2
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: "BrightScript in section \"Evaluate a flag\""
# Bucket C: no Roku BrightScript validator. See _sdk-docs-port-notes.md.
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only
---

```brightscript
while (true)

    ' do not wait forever or timers will break
    msg = wait(3000, messagePort)


    if launchDarkly.handleMessage(msg) then
        ' this message was for the client
    else
        ' handle non client messages
    end if
end while
```
