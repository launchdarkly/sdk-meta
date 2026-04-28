FROM haskell:9.6.7-bullseye

# launchdarkly-server-sdk's transitive deps need libpcre + zlib + openssl.
RUN apt-get update && apt-get install -y --no-install-recommends \
        libpcre3-dev zlib1g-dev libssl-dev \
        ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Pre-build a cabal project mirroring what the snippet's `stack new` flow
# produces, so per-validate cycles only recompile the user's Main.hs.
# The snippet's stack.yaml/package.yaml fragments aren't exercised by the
# validator (they're just rendered into ld-application's get-started
# instructions); we build the equivalent cabal project here directly.
ARG LD_HASKELL_VERSION=4.5.1
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
