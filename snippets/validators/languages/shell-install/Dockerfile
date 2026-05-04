# Build context is `validators/`. See validators/languages/python/Dockerfile
# for the rationale.
#
# Multi-toolchain image: each sdk-info install snippet uses one of npm /
# pnpm / yarn / pip / go-get. The harness sniffs the staged body and runs
# the appropriate package manager. We base on node:20 (Debian bookworm)
# and layer Python and Go on top so a single image covers every Linux
# install command we ship today.
FROM node:20-bookworm-slim

# pnpm + yarn come via corepack (node 20 ships it). Python provides the
# pip path. Go provides `go get`.
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

# Enable corepack-shipped package managers without prompting.
ENV COREPACK_ENABLE_DOWNLOAD_PROMPT=0
RUN corepack enable

WORKDIR /work

COPY shared /harness-shared
COPY languages/shell-install/harness /harness

ENTRYPOINT ["/harness/run.sh"]
