// Loads the staged HTML in headless Chromium, polls the page text for the
// EXAM-HELLO success line, and exits 0 when matched. The browser is given
// up to 30 seconds to fetch the LaunchDarkly client SDK from its CDN,
// initialize, and evaluate the flag.
const { chromium } = require('/harness/node_modules/playwright');

(async () => {
  const entrypoint = process.env.SNIPPET_ENTRYPOINT || 'index.html';
  const url = `file:///snippet/${entrypoint}`;
  const successRe = /feature flag evaluates to true/i;

  const browser = await chromium.launch();
  const context = await browser.newContext();
  const page = await context.newPage();

  // Mirror page console messages to validator stdout so a snippet that
  // fails to init is debuggable from the run log.
  page.on('console', msg => console.log('[browser]', msg.text()));
  page.on('pageerror', err => console.log('[browser:error]', err.message));

  await page.goto(url);

  const deadline = Date.now() + 30_000;
  while (Date.now() < deadline) {
    const text = (await page.textContent('body')) || '';
    if (successRe.test(text)) {
      console.log(text.trim());
      console.log('validator: ok');
      await browser.close();
      process.exit(0);
    }
    await page.waitForTimeout(500);
  }

  const finalText = (await page.textContent('body')) || '';
  console.error('validator: did not see expected line: feature flag evaluates to true');
  console.error('--- final body text ---');
  console.error(finalText.trim());
  await browser.close();
  process.exit(1);
})().catch(async (err) => {
  console.error('validator: harness error:', err.message);
  process.exit(1);
});
