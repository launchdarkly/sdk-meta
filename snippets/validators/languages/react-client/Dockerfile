FROM mcr.microsoft.com/playwright:v1.59.1-noble

# Pre-bake a minimal Vite + React + TypeScript project with the LD React
# SDK installed. Per-validate cycles drop the snippet's entrypoint
# (src/index.tsx for legacy, src/main.tsx for createApp) plus its App.tsx
# companion into src/ and re-run vite build.
ARG LD_REACT_SDK_VERSION=3.9.0
ARG REACT_VERSION=19.2.5
ARG VITE_VERSION=8.0.10
ARG VITE_PLUGIN_REACT_VERSION=6.0.1

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
    "launchdarkly-react-client-sdk": "%s",\n\
    "react": "%s",\n\
    "react-dom": "%s"\n\
  },\n\
  "devDependencies": {\n\
    "@types/react": "19.2.5",\n\
    "@types/react-dom": "19.2.3",\n\
    "@vitejs/plugin-react": "%s",\n\
    "typescript": "5.6.3",\n\
    "vite": "%s"\n\
  }\n\
}\n' \
    "${LD_REACT_SDK_VERSION}" "${REACT_VERSION}" "${REACT_VERSION}" \
    "${VITE_PLUGIN_REACT_VERSION}" "${VITE_VERSION}" > package.json

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

RUN cd /harness && npm install --silent --no-audit --no-fund --no-progress playwright@1.59.1

ENTRYPOINT ["/harness/run.sh"]
