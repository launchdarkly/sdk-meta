FROM mcr.microsoft.com/playwright:v1.59.1-noble

# Flutter SDK + the Linux desktop build toolchain. The snippets are
# validated with `flutter build linux --debug` (a real Dart front-end
# compile that stops before dart2js/AOT — see harness/run.sh), so the
# image needs the desktop build deps: clang + cmake + ninja + GTK headers.
ARG FLUTTER_VERSION=3.27.4

RUN apt-get update && apt-get install -y --no-install-recommends \
        curl xz-utils git unzip ca-certificates \
        # Linux desktop build toolchain (flutter build linux):
        clang cmake ninja-build pkg-config libgtk-3-dev liblzma-dev \
    && rm -rf /var/lib/apt/lists/*

# Flutter wants a non-root user by default; per their warning we set
# the runner ENV. We're inside an isolated build container so root is fine.
ENV PUB_CACHE=/opt/pub-cache
ENV FLUTTER_HOME=/opt/flutter
ENV PATH="${FLUTTER_HOME}/bin:${PATH}"

RUN curl -fsSL "https://storage.googleapis.com/flutter_infra_release/releases/stable/linux/flutter_linux_${FLUTTER_VERSION}-stable.tar.xz" \
        | tar -xJ -C /opt \
    && git config --global --add safe.directory /opt/flutter \
    && flutter config --no-analytics --enable-linux-desktop

# Pre-bake a hello-flutter linux desktop project with the SDK + provider
# deps already pinned. Per-validate cycles overwrite lib/main.dart and
# re-run `flutter build linux --debug` against the warmed pub cache + dart
# kernel cache + the pre-built desktop embedder.
ARG LD_FLUTTER_SDK_VERSION=4.20.0
ARG PROVIDER_VERSION=6.1.2

RUN flutter create --platforms=linux --org com.launchdarkly /opt/hello_flutter
WORKDIR /opt/hello_flutter

# Replace the default deps with our pinned set so the lockfile is
# stable and the warmed kernel cache below applies to every per-validate
# build.
RUN sed -i "/^dependencies:/,/^[a-z]/{/^  cupertino_icons/d}" pubspec.yaml \
    && sed -i "/^dependencies:/a\\  launchdarkly_flutter_client_sdk: ${LD_FLUTTER_SDK_VERSION}\\n  provider: ${PROVIDER_VERSION}" pubspec.yaml \
    && flutter pub get

# Pre-warm the linux debug build so the first per-validate is incremental:
# this bakes the CMake configuration and the compiled desktop embedder into
# the image, leaving only the Dart kernel recompile per snippet.
RUN flutter build linux --debug

WORKDIR /work

COPY shared /harness-shared
COPY languages/flutter-client/harness /harness

ENTRYPOINT ["/harness/run.sh"]
