# Build context is `validators/`. See validators/languages/python/Dockerfile
# for the rationale.
#
# Parse-only validator for Roku BrightScript snippets. There's no
# vendor-supplied BrightScript compiler we can run in CI; Roku's IDE
# (Roku Studio) is the canonical compile path and ships only as a
# GUI. The RokuCommunity `brighterscript` toolchain provides a
# Node-based parser/lexer that fills the gap for syntax validation.
# It accepts both classic BrightScript (`.brs`) and the BrighterScript
# superset.
FROM node:20-bookworm-slim

# Pin to the latest stable release so the validator is reproducible;
# bump when we want newer language-feature support.
ARG BRIGHTERSCRIPT_VERSION=0.72.4
RUN npm install -g brighterscript@${BRIGHTERSCRIPT_VERSION}

# Expose the global node_modules dir on NODE_PATH so the harness's
# `require('brighterscript')` resolves without a brittle absolute path.
ENV NODE_PATH=/usr/local/lib/node_modules

# The Lexer + Parser are invoked from a tiny Node wrapper in the
# harness. No project file or pre-baked workspace needed — each
# snippet is its own self-contained `.brs` file.

WORKDIR /work

COPY shared /harness-shared
COPY languages/roku-client/harness /harness

ENTRYPOINT ["/harness/run.sh"]
