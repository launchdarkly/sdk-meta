# Toolchain must satisfy the SDK's MSRV: cargo's MSRV-aware resolver
# picks the newest launchdarkly-server-sdk whose rust-version the
# toolchain supports, and the 3.x line (flush_blocking, etc.)
# requires 1.94.
FROM rust:1.94

WORKDIR /work

COPY shared /harness-shared
COPY languages/rust/harness /harness

ENTRYPOINT ["/harness/run.sh"]
