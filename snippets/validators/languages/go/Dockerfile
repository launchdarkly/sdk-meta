# Build context is `validators/`. See validators/languages/python/Dockerfile
# for the rationale.
#
# Track the go-server-sdk's stated minimum: v7.15.5 declares `go >= 1.25.0`.
# The official golang images set GOTOOLCHAIN=local, so an under-spec base
# does not silently self-upgrade -- `go mod tidy` fails outright with
# "toolchain upgrade needed". Bump this when the SDK's minimum moves.
FROM golang:1.25

WORKDIR /work

COPY shared /harness-shared
COPY languages/go/harness /harness

ENTRYPOINT ["/harness/run.sh"]
