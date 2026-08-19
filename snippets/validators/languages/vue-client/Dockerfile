ARG PLAYWRIGHT_IMAGE
FROM ${PLAYWRIGHT_IMAGE}

COPY shared/versions/npm.env /versions/npm.env
RUN . /versions/npm.env && npm install -g npm@"$NPM_VERSION" \
 && printf 'min-release-age=3\nmin-release-age-exclude[]=@launchdarkly/*\nmin-release-age-exclude[]=launchdarkly-*\n' > /root/.npmrc

# PLAYWRIGHT_IMAGE and PLAYWRIGHT_VERSION move together.
# Pre-bake a minimal Vue 3 + Vite project with the LD Vue SDK and a stable
# index.html. Per-validate cycles drop the snippet's main.js + App.vue
# into src/ and re-run vite build; the SDK + Vue + Vite are already cached.
ARG LD_VUE_SDK_VERSION
ARG LD_VUE_OBSERVABILITY_VERSION
ARG LD_VUE_SESSION_REPLAY_VERSION
ARG VUE_VERSION
ARG VITE_VERSION
ARG VITE_PLUGIN_VUE_VERSION

RUN mkdir -p /opt/hello-vue/src
WORKDIR /opt/hello-vue

RUN printf '{\n\
  "name": "hello-vue",\n\
  "private": true,\n\
  "version": "0.0.0",\n\
  "type": "module",\n\
  "scripts": {\n\
    "dev": "vite",\n\
    "build": "vite build",\n\
    "preview": "vite preview --port=4173 --host"\n\
  },\n\
  "dependencies": {\n\
    "launchdarkly-vue-client-sdk": "%s",\n\
    "@launchdarkly/observability": "%s",\n\
    "@launchdarkly/session-replay": "%s",\n\
    "vue": "%s"\n\
  },\n\
  "devDependencies": {\n\
    "@vitejs/plugin-vue": "%s",\n\
    "vite": "%s"\n\
  }\n\
}\n' "${LD_VUE_SDK_VERSION}" "${LD_VUE_OBSERVABILITY_VERSION}" "${LD_VUE_SESSION_REPLAY_VERSION}" \
     "${VUE_VERSION}" "${VITE_PLUGIN_VUE_VERSION}" "${VITE_VERSION}" > package.json

RUN printf "import { defineConfig } from 'vite'\n\
import vue from '@vitejs/plugin-vue'\n\
\n\
export default defineConfig({\n\
  plugins: [vue()],\n\
})\n" > vite.config.js

RUN printf '<!doctype html>\n\
<html lang="en">\n\
  <head><meta charset="UTF-8" /><title>hello-vue</title></head>\n\
  <body>\n\
    <div id="app"></div>\n\
    <script type="module" src="/src/main.js"></script>\n\
  </body>\n\
</html>\n' > index.html

# Placeholder src files. Replaced per-validate.
RUN printf "import { createApp } from 'vue'\nimport App from './App.vue'\ncreateApp(App).mount('#app')\n" > src/main.js
RUN printf "<template><div>placeholder</div></template>\n" > src/App.vue

RUN npm install --no-audit --no-fund --no-progress

# Pre-warm the build to seed Vite's transform cache.
RUN npm run build

WORKDIR /work

COPY shared /harness-shared
COPY languages/vue-client/harness /harness

RUN . /versions/npm.env && cd /harness && npm install --silent --no-audit --no-fund --no-progress playwright@"$PLAYWRIGHT_VERSION"

ENTRYPOINT ["/harness/run.sh"]
