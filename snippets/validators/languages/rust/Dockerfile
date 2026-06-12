# The latest launchdarkly-server-sdk crate (3.x) declares rust-version
# 1.93.0 as its MSRV; keep the toolchain ahead of it.
FROM rust:1.94

WORKDIR /work

COPY shared /harness-shared
COPY languages/rust/harness /harness

ENTRYPOINT ["/harness/run.sh"]
