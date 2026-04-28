// Drives the snippet's AppDelegate + ViewController against the live
// LaunchDarkly streaming API. The host app's AppDelegate.didFinishLaunching
// has already called LDClient.start(...) by the time this test runs.
import XCTest
import LaunchDarkly
@testable import HelloIOS

final class SnippetTest: XCTestCase {
    func testFlagEvaluatesToTrue() throws {
        let flagKey = ProcessInfo.processInfo.environment["LAUNCHDARKLY_FLAG_KEY"]
            ?? "sample-feature"

        guard let ld = LDClient.get() else {
            XCTFail("LDClient.get() returned nil — AppDelegate.setUpLDClient never ran")
            return
        }

        // The AppDelegate uses startWaitSeconds: 30, so by the time
        // didFinishLaunching returns the SDK has either initialized or
        // timed out. Poll boolVariation a few extra seconds in case the
        // first event hasn't been processed yet.
        let exp = expectation(description: "flag fetched")
        var lastResult = false
        let observer = NSObject()
        ld.observe(key: flagKey, owner: observer) { changedFlag in
            if case .bool(let b) = changedFlag.newValue, b {
                lastResult = true
                exp.fulfill()
            }
        }
        // Also seed with the current value in case streaming has already
        // delivered the flag before observe() registered.
        if ld.boolVariation(forKey: flagKey, defaultValue: false) {
            lastResult = true
            exp.fulfill()
        }
        wait(for: [exp], timeout: 30.0)

        // Drive the snippet's ViewController through the same code path
        // gonfalon shows the user. The IBOutlet is wired manually so
        // updateUi can write into a real label.
        let vc = ViewController()
        let label = UILabel()
        vc.featureFlagLabel = label
        vc.loadViewIfNeeded()
        vc.viewDidLoad()
        // viewDidLoad is async-ish — let one runloop spin so the
        // observer callback has a chance to fire.
        RunLoop.main.run(until: Date(timeIntervalSinceNow: 1.0))

        let rendered = label.text ?? ""
        print("validator: rendered=\(rendered)")
        XCTAssertTrue(lastResult, "expected flag to evaluate to true")
        XCTAssertTrue(rendered.lowercased().contains("feature flag evaluates to true"),
                      "expected canonical line in label, got: \(rendered)")
    }
}
