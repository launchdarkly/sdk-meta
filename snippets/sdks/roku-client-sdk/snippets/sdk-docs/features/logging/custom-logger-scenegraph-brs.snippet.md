---
id: roku-client-sdk/sdk-docs/features/logging/custom-logger-scenegraph-brs
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: SceneGraph API custom logger task implementation for Roku (CustomLogger.brs).
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only-toplevel

---

```brightscript
' /components/CustomLogger.brs

function init()

    m.messagePort = createObject("roMessagePort")


    m.top.observeField("log", m.messagePort)

end function

function mainThread() as Void
    while (true)

        msg = wait(0, m.messagePort)


        if type(msg) = "roSGNodeEvent" then
            if msg.getField() = "log" then
                value = msg.getData()

                print value.level value.message
            end if
        end if
    end while
end function
```
