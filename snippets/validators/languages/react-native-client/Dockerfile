FROM node:22-bookworm

# React Native's jest setup pulls in a long tail of transitive deps;
# pre-bake them so per-validate cycles only re-run jest with the
# snippet's two .tsx files.

ARG LD_RN_SDK_VERSION=10.17.2
ARG REACT_NATIVE_VERSION=0.85.2
ARG REACT_VERSION=19.2.5
ARG TESTING_LIBRARY_VERSION=13.3.3

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
    "test": "jest --colors=false"
  },
  "dependencies": {
    "@launchdarkly/react-native-client-sdk": "${LD_RN_SDK_VERSION}",
    "@react-native-async-storage/async-storage": "^1.23.1",
    "react": "${REACT_VERSION}",
    "react-native": "${REACT_NATIVE_VERSION}"
  },
  "devDependencies": {
    "@babel/core": "^7.25.2",
    "@babel/preset-env": "^7.25.3",
    "@react-native/babel-preset": "${REACT_NATIVE_VERSION}",
    "@react-native/jest-preset": "${REACT_NATIVE_VERSION}",
    "@testing-library/react-native": "${TESTING_LIBRARY_VERSION}",
    "@types/react": "${REACT_VERSION}",
    "@types/react-test-renderer": "19.1.0",
    "jest": "^29.7.0",
    "react-test-renderer": "${REACT_VERSION}",
    "ts-jest": "^29.2.5",
    "typescript": "^5.6.0"
  },
  "jest": {
    "preset": "@react-native/jest-preset",
    "transformIgnorePatterns": [
      "node_modules/(?!(react-native|@react-native|@launchdarkly)/)"
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
