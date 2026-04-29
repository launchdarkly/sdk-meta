package com.launchdarkly.hello_android

import android.widget.TextView
import androidx.test.core.app.ApplicationProvider
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.Robolectric
import org.robolectric.RobolectricTestRunner
import org.robolectric.annotation.Config

/**
 * Robolectric test that drives MainApplication + MainActivity end-to-end
 * against the real LaunchDarkly streaming API. This is the validator's
 * harness; it lives in app/src/test/ and is added to the project at
 * docker-build time. The snippet's two .kt files are dropped into
 * app/src/main/ at validate time and compiled alongside this test.
 */
@RunWith(RobolectricTestRunner::class)
@Config(sdk = [33], application = MainApplication::class)
class HelloAppTest {
    @Test
    fun flagEvaluatesToTrue() {
        // ApplicationProvider triggers MainApplication.onCreate which
        // calls LDClient.init with the mobile key the snippet baked in.
        val app = ApplicationProvider.getApplicationContext<MainApplication>()
        check(app != null) { "MainApplication context not registered" }

        // Drive the activity lifecycle. MainActivity.onCreate calls
        // boolVariation and renders into the TextView.
        val controller = Robolectric.buildActivity(MainActivity::class.java)
                .create()
                .start()
                .resume()
                .visible()
        val activity = controller.get()
        val textView = activity.findViewById<TextView>(R.id.textview)

        // Wait for the streaming SDK to fetch the flag and fire the
        // change listener, then for Robolectric's main looper to apply
        // the resulting setText. Polling with a wider deadline is
        // robust against transient network jitter.
        val deadline = System.currentTimeMillis() + 30_000
        var rendered = textView.text.toString()
        while (System.currentTimeMillis() < deadline) {
            rendered = textView.text.toString()
            if (rendered.contains("evaluates to true", ignoreCase = true)) {
                println("validator: ok")
                println(rendered)
                return
            }
            Thread.sleep(500)
            org.robolectric.shadows.ShadowLooper.runUiThreadTasksIncludingDelayedTasks()
        }
        assertTrue("did not see expected line; got: $rendered",
                rendered.contains("evaluates to true", ignoreCase = true))
    }
}
