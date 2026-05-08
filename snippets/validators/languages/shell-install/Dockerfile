# Build context is `validators/`. See validators/languages/python/Dockerfile
# for the rationale.
#
# Multi-toolchain image: each sdk-info / observability / ai-configs / sdk-docs
# install snippet uses one of npm / pnpm / yarn / pip / go-get / gem /
# cargo / composer / dotnet (CLI) / gradle. The harness sniffs the staged
# body's leading token and runs the appropriate package manager. Image is
# intentionally fat (~2 GB) — a CI-only artifact, cached per matrix cell
# after the first build, so the size pays for itself quickly.
FROM node:20-bookworm-slim

# pnpm + yarn come via corepack (node 20 ships it). Python provides the
# pip path. Go provides `go get`. Ruby provides the `gem install` path
# used by ruby-server-sdk install snippets.
#
# git is needed because `go get` resolves modules over git, and bower (if
# we ever re-enable it) clones from git remotes.
#
# ca-certificates / curl are needed for fetching the Go tarball below.
RUN apt-get update \
 && apt-get install -y --no-install-recommends \
        ca-certificates \
        curl \
        git \
        python3 \
        python3-venv \
        python3-pip \
        ruby \
        ruby-dev \
        build-essential \
        zlib1g-dev \
        libssl-dev \
        libyaml-dev \
        libffi-dev \
 && rm -rf /var/lib/apt/lists/*

# Install Go. Pinning here so the image is reproducible; bump when the
# go-server-sdk's stated minimum moves.
ENV GO_VERSION=1.24.3
RUN curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" \
        -o /tmp/go.tgz \
 && tar -C /usr/local -xzf /tmp/go.tgz \
 && rm /tmp/go.tgz
ENV PATH=/usr/local/go/bin:/root/go/bin:$PATH
ENV GOFLAGS=-mod=mod

# Enable corepack-shipped package managers without prompting. Pin pnpm and
# yarn to versions compatible with the Node 20 base above; without an
# explicit pin, corepack auto-fetches `latest` on first invocation, and pnpm
# 11+ requires Node 22.13+ (uses the `node:sqlite` builtin) which crashes
# the harness with `ERR_UNKNOWN_BUILTIN_MODULE`. pnpm 9 is the last major
# that supports Node 20; yarn 1 is the classic line that every install
# snippet uses.
ENV COREPACK_ENABLE_DOWNLOAD_PROMPT=0
RUN corepack enable \
 && corepack prepare pnpm@9 --activate \
 && corepack prepare yarn@1 --activate

# Rust toolchain for `cargo add` / `cargo init` install snippets. Use the
# rustup minimal profile to avoid pulling docs and component metadata we
# don't need at runtime.
ENV RUSTUP_HOME=/usr/local/rustup
ENV CARGO_HOME=/usr/local/cargo
ENV PATH=/usr/local/cargo/bin:$PATH
RUN curl -fsSL https://sh.rustup.rs \
        | sh -s -- -y --default-toolchain stable --profile minimal --no-modify-path

# PHP + composer.phar for `php composer.phar require …` install snippets.
# composer.phar lands at /opt/composer.phar; the harness symlinks it into
# the working dir before running the body so the body's literal
# `composer.phar` reference resolves.
RUN apt-get update \
 && apt-get install -y --no-install-recommends \
        php-cli \
        php-zip \
        php-curl \
        php-mbstring \
        unzip \
 && rm -rf /var/lib/apt/lists/*
RUN curl -fsSL https://getcomposer.org/download/latest-stable/composer.phar -o /opt/composer.phar \
 && chmod +x /opt/composer.phar

# .NET SDK for `dotnet add package` install snippets. The Microsoft apt
# feed serves a .deb that wires up the package source for us.
RUN curl -fsSL https://packages.microsoft.com/config/debian/12/packages-microsoft-prod.deb \
        -o /tmp/packages-microsoft-prod.deb \
 && dpkg -i /tmp/packages-microsoft-prod.deb \
 && rm /tmp/packages-microsoft-prod.deb \
 && apt-get update \
 && apt-get install -y --no-install-recommends dotnet-sdk-8.0 \
 && rm -rf /var/lib/apt/lists/*

# OpenJDK + Gradle for `implementation group: …` Gradle DSL install
# snippets. The harness wraps the body in a minimal build.gradle and
# runs `gradle dependencies` to confirm the artifact resolves. Pin
# Gradle to a recent version (Debian's apt package lags by years).
RUN apt-get update \
 && apt-get install -y --no-install-recommends openjdk-17-jdk-headless \
 && rm -rf /var/lib/apt/lists/*
ENV GRADLE_VERSION=8.10.2
RUN curl -fsSL "https://services.gradle.org/distributions/gradle-${GRADLE_VERSION}-bin.zip" \
        -o /tmp/gradle.zip \
 && unzip -q /tmp/gradle.zip -d /opt \
 && ln -s "/opt/gradle-${GRADLE_VERSION}/bin/gradle" /usr/local/bin/gradle \
 && rm /tmp/gradle.zip

WORKDIR /work

COPY shared /harness-shared
COPY languages/shell-install/harness /harness

ENTRYPOINT ["/harness/run.sh"]
