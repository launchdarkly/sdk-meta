FROM mcr.microsoft.com/playwright:v1.59.1-noble

# Version-pinned counterpart to the current flutter-client validator.
# Pins `launchdarkly_flutter_client_sdk` to a 2.x release so doc
# fragments under `sdk-docs/initialize-the-client-flutter-sdk-v2-x*`
# (which reference `LDConfigBuilder`, `LDClient.startWithContext` —
# names removed in 3.x+) can validate against the actual v2.x API
# surface the docs cover. The current-version flutter-client
# validator builds against 4.x and would fail these fragments.

ARG FLUTTER_VERSION=3.27.4

RUN apt-get update && apt-get install -y --no-install-recommends \
        curl xz-utils git unzip ca-certificates \
        # Linux desktop build toolchain (flutter build linux):
        clang cmake ninja-build pkg-config libgtk-3-dev liblzma-dev \
    && rm -rf /var/lib/apt/lists/*

ENV PUB_CACHE=/opt/pub-cache
ENV FLUTTER_HOME=/opt/flutter
ENV PATH="${FLUTTER_HOME}/bin:${PATH}"

RUN curl -fsSL "https://storage.googleapis.com/flutter_infra_release/releases/stable/linux/flutter_linux_${FLUTTER_VERSION}-stable.tar.xz" \
        | tar -xJ -C /opt \
    && git config --global --add safe.directory /opt/flutter \
    && flutter config --no-analytics --enable-linux-desktop

ARG LD_FLUTTER_SDK_VERSION=2.1.0
ARG PROVIDER_VERSION=6.1.2

RUN flutter create --platforms=linux --org com.launchdarkly /opt/hello_flutter
WORKDIR /opt/hello_flutter

RUN sed -i "/^dependencies:/,/^[a-z]/{/^  cupertino_icons/d}" pubspec.yaml \
    && sed -i "/^dependencies:/a\\  launchdarkly_flutter_client_sdk: ${LD_FLUTTER_SDK_VERSION}\\n  provider: ${PROVIDER_VERSION}" pubspec.yaml \
    && flutter pub get

RUN flutter build linux --debug

WORKDIR /work

COPY shared /harness-shared
COPY languages/flutter-client-v2/harness /harness

ENTRYPOINT ["/harness/run.sh"]
