ARG NODE_BOOKWORM_IMAGE
FROM ${NODE_BOOKWORM_IMAGE}

COPY shared/versions/npm.env /versions/npm.env
RUN . /versions/npm.env && npm install -g npm@"$NPM_VERSION" \
 && printf 'min-release-age=3\nmin-release-age-exclude[]=@launchdarkly/*\nmin-release-age-exclude[]=launchdarkly-*\nmin-release-age-exclude[]=highlight.run\n' > /root/.npmrc

# React Native's jest setup pulls in a long tail of transitive deps;
# pre-bake them so per-validate cycles only re-run jest with the
# snippet's two .tsx files.

ARG LD_RN_SDK_VERSION
ARG LD_OBSERVABILITY_RN_VERSION
ARG REACT_NATIVE_VERSION
ARG REACT_VERSION
ARG TESTING_LIBRARY_VERSION
ARG RN_ASYNC_STORAGE_VERSION
ARG RN_BABEL_CORE_VERSION
ARG RN_BABEL_PRESET_ENV_VERSION
ARG RN_TYPES_REACT_TEST_RENDERER_VERSION
ARG RN_JEST_VERSION
ARG RN_TS_JEST_VERSION
ARG RN_TYPESCRIPT_VERSION

WORKDIR /opt/hello-react-native

# Minimal package.json — just enough for jest + RN testing-library to
# resolve the snippet's imports. Native code paths are mocked out by
# react-native's jest preset; the SDK's streaming layer talks to LD
# through the same JS code path it would use on a real device.
RUN cat > package.json <<EOF
{
  "name": "hello-react-native",
  "version": "0.0.1",
  "private": true,
  "scripts": {
    "test": "jest --colors=false --forceExit"
  },
  "dependencies": {
    "@launchdarkly/react-native-client-sdk": "${LD_RN_SDK_VERSION}",
    "@launchdarkly/observability-react-native": "${LD_OBSERVABILITY_RN_VERSION}",
    "@react-native-async-storage/async-storage": "${RN_ASYNC_STORAGE_VERSION}",
    "react": "${REACT_VERSION}",
    "react-native": "${REACT_NATIVE_VERSION}"
  },
  "devDependencies": {
    "@babel/core": "${RN_BABEL_CORE_VERSION}",
    "@babel/preset-env": "${RN_BABEL_PRESET_ENV_VERSION}",
    "@react-native/babel-preset": "${REACT_NATIVE_VERSION}",
    "@react-native/jest-preset": "${REACT_NATIVE_VERSION}",
    "@testing-library/react-native": "${TESTING_LIBRARY_VERSION}",
    "@types/react": "${REACT_VERSION}",
    "@types/react-test-renderer": "${RN_TYPES_REACT_TEST_RENDERER_VERSION}",
    "jest": "${RN_JEST_VERSION}",
    "react-test-renderer": "${REACT_VERSION}",
    "ts-jest": "${RN_TS_JEST_VERSION}",
    "typescript": "${RN_TYPESCRIPT_VERSION}"
  },
  "jest": {
    "preset": "@react-native/jest-preset",
    "transformIgnorePatterns": [
      "node_modules/(?!(react-native|react-native-url-polyfill|@react-native|@launchdarkly)/)"
    ],
    "moduleFileExtensions": ["ts", "tsx", "js", "jsx", "json"],
    "testMatch": ["<rootDir>/__tests__/**/*.test.tsx"],
    "setupFiles": ["<rootDir>/jest.setup.js"]
  }
}
EOF

RUN cat > babel.config.js <<'EOF'
module.exports = {
  presets: ['@react-native/babel-preset'],
};
EOF

# jest setup: react-native's preset leaves AppState.currentState as
# undefined, which the LD SDK's RNStateDetector translates to
# `Background`. With automaticBackgroundHandling=true (the default) and
# runInBackground=false (the default) the SDK then immediately enters
# offline mode, so the streaming connection never opens. Pin
# currentState to 'active' so the SDK treats us as foregrounded.
RUN cat > jest.setup.js <<'EOF'
const RN = require('react-native');
if (RN.AppState) {
  RN.AppState.currentState = 'active';
}
EOF

RUN cat > tsconfig.json <<'EOF'
{
  "compilerOptions": {
    "target": "ES2020",
    "module": "ESNext",
    "moduleResolution": "node",
    "jsx": "react-native",
    "esModuleInterop": true,
    "skipLibCheck": true,
    "strict": false
  }
}
EOF

# Placeholder snippet files — replaced per-validate.
RUN mkdir -p src __tests__ \
    && cat > App.tsx <<'EOF'
import React from 'react';
import {Text, View} from 'react-native';
const App = () => <View><Text>placeholder</Text></View>;
export default App;
EOF

RUN cat > src/welcome.tsx <<'EOF'
import React from 'react';
import {Text, View} from 'react-native';
export default function Welcome() {
  return <View><Text>placeholder</Text></View>;
}
EOF

WORKDIR /opt/hello-react-native

RUN npm install --no-audit --no-fund --no-progress

# Drop a jest test that renders App and waits for the canonical line.
COPY languages/react-native-client/test/App.test.tsx __tests__/App.test.tsx

WORKDIR /work

COPY shared /harness-shared
COPY languages/react-native-client/harness /harness

ENTRYPOINT ["/harness/run.sh"]
