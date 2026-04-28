// Loads the locally-served Flutter web bundle in headless Chromium and
// polls the rendered Flutter canvas for the EXAM-HELLO success line.
//
// Flutter web renders text into <flt-semantics> elements (semantics
// tree) once the page has bootstrapped. Some Flutter releases gate the
// semantics tree behind explicit activation; for those we fall back to
// the raw <flutter-view> DOM, which still contains the rendered text in
// HTML mode. We probe both.
const { chromium } = require('/harness/node_modules/playwright');

(async () => {
  const url = process.env.FLUTTER_PREVIEW_URL || 'http://localhost:4173';
  const successRe = /feature flag evaluates to true/i;

  const browser = await chromium.launch();
  const context = await browser.newContext();
  const page = await context.newPage();

  page.on('console', msg => console.log('[browser]', msg.text()));
  page.on('pageerror', err => console.log('[browser:error]', err.message));

  for (let i = 0; i < 25; i++) {
    try {
      await page.goto(url);
      break;
    } catch {
      await new Promise(r => setTimeout(r, 200));
    }
  }

  // Force semantic tree on so the Text widget's content lands in the DOM.
  await page.evaluate(() => {
    if (window.flutterSemanticsTree) return;
    const enable = () => {
      try {
        const placeholder = document.querySelector('flt-semantics-placeholder');
        if (placeholder) placeholder.click();
      } catch {}
    };
    enable();
    setTimeout(enable, 1000);
  });

  const deadline = Date.now() + 60_000;
  while (Date.now() < deadline) {
    const text = (await page.textContent('body')) || '';
    if (successRe.test(text)) {
      console.log(text.replace(/\s+/g, ' ').trim());
      console.log('validator: ok');
      await browser.close();
      process.exit(0);
    }
    await page.waitForTimeout(500);
  }

  const finalText = (await page.textContent('body')) || '';
  console.error('validator: did not see expected line: feature flag evaluates to true');
  console.error('--- final body text ---');
  console.error(finalText.replace(/\s+/g, ' ').trim());
  await browser.close();
  process.exit(1);
})().catch((err) => {
  console.error('validator: harness error:', err.message);
  process.exit(1);
});
