// Loads the locally-served Vue app in headless Chromium, polls the page
// text for the EXAM-HELLO success line, and exits 0 when matched.
const { chromium } = require('/harness/node_modules/playwright');

(async () => {
  const url = process.env.VUE_PREVIEW_URL || 'http://localhost:4173';
  const successRe = /feature flag evaluates to true/i;

  const browser = await chromium.launch();
  const context = await browser.newContext();
  const page = await context.newPage();

  page.on('console', msg => console.log('[browser]', msg.text()));
  page.on('pageerror', err => console.log('[browser:error]', err.message));

  // Vite preview may take ~250ms to be ready after spawn. Tolerate ECONNREFUSED.
  for (let i = 0; i < 25; i++) {
    try {
      await page.goto(url);
      break;
    } catch {
      await new Promise(r => setTimeout(r, 200));
    }
  }

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
})().catch((err) => {
  console.error('validator: harness error:', err.message);
  process.exit(1);
});
