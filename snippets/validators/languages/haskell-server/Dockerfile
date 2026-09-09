ARG DEBIAN_BOOKWORM_SLIM_IMAGE
FROM ${DEBIAN_BOOKWORM_SLIM_IMAGE}

# launchdarkly-server-sdk-redis-hedis pins bytestring < 0.12, so this validator
# needs GHC 9.6, whose only official haskell image variant is EOL bullseye;
# ghcup installs the same toolchain onto a supported base instead.
ARG GHC_VERSION
ARG CABAL_VERSION
ENV PATH=/root/.ghcup/bin:$PATH

# launchdarkly-server-sdk's transitive deps need libpcre + zlib + openssl;
# the rest is what ghcup's GHC bindist requires.
RUN apt-get update && apt-get install -y --no-install-recommends \
        libpcre3-dev zlib1g-dev libssl-dev \
        ca-certificates curl gcc g++ make libgmp-dev libnuma-dev \
        libtinfo-dev xz-utils \
    && rm -rf /var/lib/apt/lists/*

RUN curl -sSLf https://get-ghcup.haskell.org | \
        BOOTSTRAP_HASKELL_NONINTERACTIVE=1 \
        BOOTSTRAP_HASKELL_MINIMAL=1 sh \
    && ghcup install ghc "${GHC_VERSION}" --set \
    && ghcup install cabal "${CABAL_VERSION}" --set

# Pre-build a cabal project mirroring what the snippet's `stack new` flow
# produces, so per-validate cycles only recompile the user's Main.hs.
# The snippet's stack.yaml/package.yaml fragments aren't exercised by the
# validator (they're just rendered into ld-application's get-started
# instructions); we build the equivalent cabal project here directly.
ARG LD_HASKELL_VERSION
RUN mkdir -p /opt/hello-haskell/app
WORKDIR /opt/hello-haskell

RUN printf 'cabal-version: 2.2\n\
name: hello-haskell\n\
version: 0.1.0.0\n\
\n\
executable hello-haskell-exe\n\
  main-is: Main.hs\n\
  hs-source-dirs: app\n\
  build-depends:\n\
    base >= 4.7 && < 5,\n\
    text,\n\
    hedis,\n\
    launchdarkly-server-sdk-redis-hedis,\n\
    containers,\n\
    aeson,\n\
    launchdarkly-server-sdk == %s\n\
  default-language: Haskell2010\n\
  ghc-options: -threaded -rtsopts -with-rtsopts=-N\n' \
        "${LD_HASKELL_VERSION}" > hello-haskell.cabal

# Placeholder Main.hs — replaced per-validate. We need *some* file here
# so `cabal build` succeeds and primes the dependency cache.
RUN printf 'module Main where\nmain :: IO ()\nmain = putStrLn "placeholder"\n' > app/Main.hs

RUN cabal update && cabal build --only-dependencies && cabal build

WORKDIR /work

COPY shared /harness-shared
COPY languages/haskell-server/harness /harness

ENTRYPOINT ["/harness/run.sh"]
