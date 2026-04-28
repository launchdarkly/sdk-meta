---
id: roku-client-sdk/getting-started/main-brs
sdk: roku-client-sdk
kind: hello-world
lang: brightscript
file: source/main.brs
description: Roku scene runner that boots the AppScene.
ld-application:
  slot: main-brs
---

Create a file with a basic scene runner named `source/main.brs` and add the following code:

```brightscript
sub main(params as object)
  screen = createObject("roSGScreen")
  messagePort = createObject("roMessagePort")
  screen.setMessagePort(messagePort)

  scene = screen.CreateScene("AppScene")

  screen.show()

  while (true)
      msg = wait(2500, messagePort)

      if type(msg) = "roSGScreenEvent"
          if msg.isScreenClosed() then
              return
          end if
      end if
  end while
end sub
```
