ARG PLAYWRIGHT_IMAGE
FROM ${PLAYWRIGHT_IMAGE}

COPY shared/versions/npm.env /versions/npm.env
RUN . /versions/npm.env && npm install -g npm@"$NPM_VERSION" \
 && printf 'min-release-age=3\nmin-release-age-exclude[]=@launchdarkly/*\nmin-release-age-exclude[]=launchdarkly-*\nmin-release-age-exclude[]=highlight.run\n' > /root/.npmrc

# Pre-bake a minimal Vite + React + TypeScript project with the LD React
# SDK installed. Per-validate cycles drop the snippet's entrypoint
# (src/index.tsx for legacy, src/main.tsx for createApp) plus its App.tsx
# companion into src/ and re-run vite build.
ARG LD_REACT_SDK_VERSION
ARG LD_REACT_LEGACY_SDK_VERSION
ARG LD_REACT_OBSERVABILITY_VERSION
ARG LD_REACT_SESSION_REPLAY_VERSION
ARG REACT_VERSION
ARG VITE_VERSION
ARG VITE_PLUGIN_REACT_VERSION
ARG REACT_TYPES_VERSION
ARG REACT_DOM_TYPES_VERSION
ARG REACT_TYPESCRIPT_VERSION

RUN mkdir -p /opt/hello-react/src
WORKDIR /opt/hello-react

RUN printf '{\n\
  "name": "hello-react",\n\
  "private": true,\n\
  "version": "0.0.0",\n\
  "type": "module",\n\
  "scripts": {\n\
    "build": "vite build",\n\
    "preview": "vite preview --port=4173 --host"\n\
  },\n\
  "dependencies": {\n\
    "@launchdarkly/react-sdk": "%s",\n\
    "launchdarkly-react-client-sdk": "%s",\n\
    "@launchdarkly/observability": "%s",\n\
    "@launchdarkly/session-replay": "%s",\n\
    "react": "%s",\n\
    "react-dom": "%s"\n\
  },\n\
  "devDependencies": {\n\
    "@types/react": "%s",\n\
    "@types/react-dom": "%s",\n\
    "@vitejs/plugin-react": "%s",\n\
    "typescript": "%s",\n\
    "vite": "%s"\n\
  }\n\
}\n' \
    "${LD_REACT_SDK_VERSION}" "${LD_REACT_LEGACY_SDK_VERSION}" \
    "${LD_REACT_OBSERVABILITY_VERSION}" "${LD_REACT_SESSION_REPLAY_VERSION}" \
    "${REACT_VERSION}" "${REACT_VERSION}" \
    "${REACT_TYPES_VERSION}" "${REACT_DOM_TYPES_VERSION}" \
    "${VITE_PLUGIN_REACT_VERSION}" "${REACT_TYPESCRIPT_VERSION}" \
    "${VITE_VERSION}" > package.json

RUN printf "import { defineConfig } from 'vite'\n\
import react from '@vitejs/plugin-react'\n\
\n\
export default defineConfig({\n\
  plugins: [react()],\n\
})\n" > vite.config.js

RUN printf '{\n\
  "compilerOptions": {\n\
    "target": "ES2020",\n\
    "lib": ["ES2020", "DOM", "DOM.Iterable"],\n\
    "module": "ESNext",\n\
    "moduleResolution": "bundler",\n\
    "jsx": "react-jsx",\n\
    "strict": false,\n\
    "isolatedModules": true,\n\
    "noEmit": true,\n\
    "skipLibCheck": true\n\
  },\n\
  "include": ["src"]\n\
}\n' > tsconfig.json

# Empty stylesheet — both legacy and createApp variants reference one.
RUN touch src/index.css src/App.css

# Placeholder entrypoint + App. Replaced per-validate. The harness rewrites
# index.html to point at whichever entrypoint the snippet declared.
RUN printf "import { createRoot } from 'react-dom/client';\nimport App from './App';\ncreateRoot(document.getElementById('root')).render(<App />);\n" > src/main.tsx
RUN printf "export default function App() { return <div>placeholder</div>; }\n" > src/App.tsx

# Stub multi-environment module so sdk-docs parse-only snippets that
# `import { ProdLDContext, StagingLDContext } from './environments'`
# resolve.
RUN printf "import { initLDReactContext } from '@launchdarkly/react-sdk';\nexport const ProdLDContext = initLDReactContext();\nexport const StagingLDContext = initLDReactContext();\n" > src/environments.ts

# Stub feedback-helper module so sdk-docs parse-only snippets that
# `import { sendFeedback } from './sendFeedback'` resolve.
RUN printf "export type LDFeedbackSentiment = 'positive' | 'neutral' | 'negative';\nexport function sendFeedback(..._args: unknown[]): void {}\n" > src/sendFeedback.ts

RUN printf '<!doctype html>\n\
<html lang="en">\n\
  <head><meta charset="UTF-8" /><title>hello-react</title></head>\n\
  <body>\n\
    <div id="root"></div>\n\
    <script type="module" src="/src/main.tsx"></script>\n\
  </body>\n\
</html>\n' > index.html

RUN npm install --no-audit --no-fund --no-progress

# Pre-warm the build to seed Vite's transform cache + esbuild prebundle.
RUN npm run build

WORKDIR /work

COPY shared /harness-shared
COPY languages/react-client/harness /harness

RUN . /versions/npm.env && cd /harness && npm install --silent --no-audit --no-fund --no-progress playwright@"$PLAYWRIGHT_VERSION"

ENTRYPOINT ["/harness/run.sh"]
