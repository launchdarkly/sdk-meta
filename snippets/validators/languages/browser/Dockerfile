# Playwright's official image bundles Chromium + headless Chrome + Node +
# Playwright. The npm install pins the matching JS-module version so the
# bundled browser binary matches the API we call against.
FROM mcr.microsoft.com/playwright:v1.59.1-noble

WORKDIR /work

COPY shared /harness-shared
COPY languages/browser/harness /harness

RUN cd /harness && npm install --silent --no-audit --no-fund --no-progress playwright@1.59.1

ENTRYPOINT ["/harness/run.sh"]
