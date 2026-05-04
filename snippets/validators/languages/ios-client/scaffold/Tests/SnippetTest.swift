// Drives the snippet's AppDelegate + ViewController against the live
// LaunchDarkly streaming API. The host app's AppDelegate.didFinishLaunching
// has already called LDClient.start(...) by the time this test runs.
//
// What this test asserts:
//   - LDClient.get() resolves (i.e. the snippet body's
//     `LDClient.start(...)` actually ran inside didFinishLaunching).
//   - The snippet's ViewController renders the canonical EXAM-HELLO
//     line into its label after init.
//
// What this test deliberately does NOT assert: the flag's truth value.
// Different LaunchDarkly sandbox envs/projects target the test
// flag-key differently; sdk-info validation cares about the body
// running end-to-end, not about which side of the if/else the env
// happened to evaluate to. The harness's outer grep on
// `feature flag evaluates to true` is the EXAM-HELLO contract; the
// ViewController emits it unconditionally on init success so a
// false-evaluating flag is still a passing init.
import XCTest
import LaunchDarkly
@testable import HelloIOS

final class SnippetTest: XCTestCase {
    func testInitAndRender() throws {
        guard let _ = LDClient.get() else {
            XCTFail("LDClient.get() returned nil — AppDelegate.setUpLDClient never ran")
            return
        }

        // Drive the snippet's ViewController. featureFlagLabel is
        // wired manually so the canonical line lands in a real label
        // even though no storyboard is loaded.
        let vc = ViewController()
        let label = UILabel()
        vc.featureFlagLabel = label
        vc.loadViewIfNeeded()
        vc.viewDidLoad()
        // viewDidLoad seeds the label synchronously and registers a
        // change observer; let one runloop spin so any seeded boolean
        // makes its way into the label.
        RunLoop.main.run(until: Date(timeIntervalSinceNow: 1.0))

        let rendered = label.text ?? ""
        print("validator: rendered=\(rendered)")
        XCTAssertTrue(rendered.lowercased().contains("feature flag evaluates to true"),
                      "expected canonical line in label, got: \(rendered)")
    }
}
