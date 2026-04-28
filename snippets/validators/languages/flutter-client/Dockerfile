FROM mcr.microsoft.com/playwright:v1.59.1-noble

# Flutter SDK + transitive build dependencies. The snippet's main.dart
# uses `provider` and Material widgets, which compile cleanly to web
# (no native plugins). Web target keeps the validator container in the
# Linux/x86 lane — no emulators or simulators required.
ARG FLUTTER_VERSION=3.27.4

RUN apt-get update && apt-get install -y --no-install-recommends \
        curl xz-utils git unzip ca-certificates \
        # Flutter doctor dependencies for web target:
        libgtk-3-0 \
    && rm -rf /var/lib/apt/lists/*

# Flutter wants a non-root user by default; per their warning we set
# the runner ENV. We're inside an isolated build container so root is fine.
ENV PUB_CACHE=/opt/pub-cache
ENV FLUTTER_HOME=/opt/flutter
ENV PATH="${FLUTTER_HOME}/bin:${PATH}"

RUN curl -fsSL "https://storage.googleapis.com/flutter_infra_release/releases/stable/linux/flutter_linux_${FLUTTER_VERSION}-stable.tar.xz" \
        | tar -xJ -C /opt \
    && git config --global --add safe.directory /opt/flutter \
    && flutter config --no-analytics --enable-web

# Pre-bake a hello-flutter web project mirroring what `flutter create`
# produces, with the SDK + provider deps already pinned. Per-validate
# cycles overwrite lib/main.dart and re-run `flutter build web` against
# the warmed pub cache + dart kernel cache.
ARG LD_FLUTTER_SDK_VERSION=4.16.0
ARG PROVIDER_VERSION=6.1.2

RUN flutter create --platforms=web --org com.launchdarkly /opt/hello_flutter
WORKDIR /opt/hello_flutter

# Replace the default deps with our pinned set so the lockfile is
# stable and the warmed kernel cache below applies to every per-validate
# build.
RUN sed -i "/^dependencies:/,/^[a-z]/{/^  cupertino_icons/d}" pubspec.yaml \
    && sed -i "/^dependencies:/a\\  launchdarkly_flutter_client_sdk: ${LD_FLUTTER_SDK_VERSION}\\n  provider: ${PROVIDER_VERSION}" pubspec.yaml \
    && flutter pub get

# Pre-warm the web build so first per-validate is incremental.
RUN flutter build web --release --no-pub

WORKDIR /work

COPY shared /harness-shared
COPY languages/flutter-client/harness /harness

RUN cd /harness && npm install --silent --no-audit --no-fund --no-progress playwright@1.59.1

ENTRYPOINT ["/harness/run.sh"]
