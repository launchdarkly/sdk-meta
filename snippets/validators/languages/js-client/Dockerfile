FROM mcr.microsoft.com/playwright:v1.59.1-noble

# Pre-bake a minimal tsdown project with the LD JavaScript browser SDK
# installed. Per-validate cycles drop src/app.ts and index.html from the
# snippet into the project and re-run tsdown.
ARG LD_JS_CLIENT_SDK_VERSION=4.6.3
ARG LD_JS_CLIENT_LEGACY_SDK_VERSION=3.8.0
ARG LD_OBSERVABILITY_VERSION=1.1.8
ARG LD_SESSION_REPLAY_VERSION=1.1.8
ARG TSDOWN_VERSION=0.17.0-beta.4
ARG TYPESCRIPT_VERSION=5.9.3

RUN mkdir -p /opt/hello-js/src
WORKDIR /opt/hello-js

RUN printf '{\n\
  "name": "hello-js",\n\
  "private": true,\n\
  "version": "0.0.0",\n\
  "type": "module",\n\
  "scripts": {\n\
    "build": "tsdown"\n\
  },\n\
  "dependencies": {\n\
    "@launchdarkly/js-client-sdk": "%s",\n\
    "launchdarkly-js-client-sdk": "%s",\n\
    "@launchdarkly/observability": "%s",\n\
    "@launchdarkly/session-replay": "%s"\n\
  },\n\
  "devDependencies": {\n\
    "tsdown": "%s",\n\
    "typescript": "%s"\n\
  }\n\
}\n' \
    "${LD_JS_CLIENT_SDK_VERSION}" "${LD_JS_CLIENT_LEGACY_SDK_VERSION}" \
    "${LD_OBSERVABILITY_VERSION}" "${LD_SESSION_REPLAY_VERSION}" \
    "${TSDOWN_VERSION}" "${TYPESCRIPT_VERSION}" > package.json

RUN printf "import { defineConfig } from 'tsdown';\n\
\n\
export default defineConfig({\n\
  entry: 'src/app.ts',\n\
  platform: 'browser',\n\
  outDir: 'dist',\n\
  format: 'iife',\n\
  outputOptions: { entryFileNames: 'app.js' },\n\
  noExternal: [/.*/],\n\
});\n" > tsdown.config.ts

RUN printf '{\n\
  "compilerOptions": {\n\
    "target": "ES2020",\n\
    "lib": ["ES2020", "DOM", "DOM.Iterable"],\n\
    "module": "ESNext",\n\
    "moduleResolution": "bundler",\n\
    "strict": false,\n\
    "isolatedModules": true,\n\
    "noEmit": true,\n\
    "skipLibCheck": true\n\
  },\n\
  "include": ["src"]\n\
}\n' > tsconfig.json

# Placeholder entrypoint + index.html. Replaced per-validate.
RUN printf "console.log('placeholder');\n" > src/app.ts
RUN printf '<!DOCTYPE html>\n\
<html lang="en">\n\
  <head>\n\
    <meta charset="utf-8">\n\
    <title>hello-js</title>\n\
    <script src="./dist/app.js" defer></script>\n\
  </head>\n\
  <body></body>\n\
</html>\n' > index.html

RUN npm install --no-audit --no-fund --no-progress

# Pre-warm the build to seed tsdown's transform cache.
RUN npm run build

WORKDIR /work

COPY shared /harness-shared
COPY languages/js-client/harness /harness

RUN cd /harness && npm install --silent --no-audit --no-fund --no-progress playwright@1.59.1

ENTRYPOINT ["/harness/run.sh"]
