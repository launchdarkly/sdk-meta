FROM mcr.microsoft.com/playwright:v1.59.1-noble

# Version-pinned counterpart to the current flutter-client validator.
# Pins `launchdarkly_flutter_client_sdk` to a 3.x release so v3.x doc
# fragments (which reference `LDConfigBuilder(key, autoEnvAttributes)`,
# `LDClient.start(config, context)` — the v3-shape init signature
# distinct from both v2.x and v4.x) validate against the actual v3.x
# SDK surface. The current-version flutter-client validator builds
# against 4.x and rejects v3-era calls.

ARG FLUTTER_VERSION=3.27.4

RUN apt-get update && apt-get install -y --no-install-recommends \
        curl xz-utils git unzip ca-certificates \
        libgtk-3-0 \
    && rm -rf /var/lib/apt/lists/*

ENV PUB_CACHE=/opt/pub-cache
ENV FLUTTER_HOME=/opt/flutter
ENV PATH="${FLUTTER_HOME}/bin:${PATH}"

RUN curl -fsSL "https://storage.googleapis.com/flutter_infra_release/releases/stable/linux/flutter_linux_${FLUTTER_VERSION}-stable.tar.xz" \
        | tar -xJ -C /opt \
    && git config --global --add safe.directory /opt/flutter \
    && flutter config --no-analytics --enable-web

ARG LD_FLUTTER_SDK_VERSION=3.1.0
ARG PROVIDER_VERSION=6.1.2

RUN flutter create --platforms=web --org com.launchdarkly /opt/hello_flutter
WORKDIR /opt/hello_flutter

RUN sed -i "/^dependencies:/,/^[a-z]/{/^  cupertino_icons/d}" pubspec.yaml \
    && sed -i "/^dependencies:/a\\  launchdarkly_flutter_client_sdk: ${LD_FLUTTER_SDK_VERSION}\\n  provider: ${PROVIDER_VERSION}" pubspec.yaml \
    && flutter pub get

RUN flutter build web --release --no-pub

WORKDIR /work

COPY shared /harness-shared
COPY languages/flutter-client-v3/harness /harness

RUN cd /harness && npm install --silent --no-audit --no-fund --no-progress playwright@1.59.1

ENTRYPOINT ["/harness/run.sh"]
