---
id: roku-client-sdk/getting-started/app-scene-brs
sdk: roku-client-sdk
kind: hello-world
lang: brightscript
file: components/AppScene.brs
description: Scene-side logic that initializes the SDK and renders the flag value.
inputs:
  mobileKey:
    type: mobile-key
    description: Mobile key baked into the rendered source. Roku snippets carry no automated validation — see the comment below the frontmatter.
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source.
ld-application:
  slot: app-scene-brs
# validation: none
#
# Roku validation requires a Roku device or the proprietary BrightScript
# simulator (no public CI runtime exists). The snippet is rendered into
# gonfalon and is reviewed manually against a real Roku device when
# changed; the `version-staleness.yml` sweep still tracks the upstream
# package.zip release. See sdk-snippet-design.md §"Validator architecture"
# for the policy.
---

In `components/AppScene.brs` add the following code:

```brightscript
function onFeatureChange() as Void
    featureFlagKey = "{{ featureKey }}"

    value = m.ld.variation(featureFlagKey, false)

    if value then
      m.top.backgroundColor = &h00844BFF
      m.evaluation.text = "The " + featureFlagKey + " feature flag evaluates to true"
    else
      m.top.backgroundColor = &h373841FF
      m.evaluation.text = "The " + featureFlagKey + " feature flag evaluates to false"
    end if
end function

function onStatusChange() as Void
    if m.ld.status.getStatus() = m.ld.status.map.initialized
      m.status.text = "SDK successfully initialized"
    else
      m.status.text = "SDK failed to initialize. Please check your internet connection and SDK credential for any typo."
    end if
end function

function init() as Void
    mobileKey = "{{ mobileKey }}"

    launchDarklyNode = m.top.findNode("launchDarkly")
    launchDarklyNode.observeField("flags", "onFeatureChange")
    launchDarklyNode.observeField("status", "onStatusChange")

    config = LaunchDarklyConfig(mobileKey, launchDarklyNode)

    ' Set up the user-kind context properties. This context should appear on
    ' your LaunchDarkly contexts dashboard soon after you run the demo.
    context = LaunchDarklyCreateContext({kind: "user", key: "example-user-key", name: "Sandy"})
    LaunchDarklySGInit(config, context)

    m.ld = LaunchDarklySG(launchDarklyNode)

    m.evaluation = m.top.findNode("evaluation")
    m.evaluation.font.size=40
    m.evaluation.color="0xFFFFFFFF"

    m.status = m.top.findNode("status")
    m.status.font.size=20
    m.status.color="0xFFFFFFFF"

    m.top.backgroundColor = &h373841FF
    m.top.backgroundUri = ""

    onFeatureChange()
end function
```
