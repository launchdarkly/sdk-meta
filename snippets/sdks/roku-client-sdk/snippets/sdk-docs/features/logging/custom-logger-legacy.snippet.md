---
id: roku-client-sdk/sdk-docs/features/logging/custom-logger-legacy
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: Legacy API custom logger example for Roku.
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only

---

```brightscript
function CustomLogger() as Object
    return {
        log: function(level as Integer, message as String)

            print level message
        end function
    }
end function

config.setLogger(CustomLogger())

```
