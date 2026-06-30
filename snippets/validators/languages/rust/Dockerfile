# Toolchain must satisfy the SDK's MSRV: cargo's MSRV-aware resolver
# picks the newest launchdarkly-server-sdk whose rust-version the
# toolchain supports, and the 3.x line (flush_blocking, etc.)
# requires 1.94.
FROM rust:1.94

# Pre-bake a warm cargo project: the SDK + tokio + futures + transport
# dependency tree is added and compiled once at image-build time. Per-validate
# cycles overwrite src/main.rs and `cargo run`, which recompiles only the
# binary crate and links against the already-cached dependency objects — the
# whole point of batch mode is to pay this dependency compile once per job
# instead of once per snippet. The crate registry index + downloaded sources
# are baked in too, so per-validate runs don't re-fetch.
#
# These `cargo add` lines mirror the set the harness assembles at runtime
# (launchdarkly-server-sdk, tokio[rt,macros], futures,
# launchdarkly-sdk-transport[hyper]); keep them in sync so the warmed cache
# matches what the default (unpinned) batch group builds.
RUN cargo new --quiet --bin /opt/hello-rust
WORKDIR /opt/hello-rust
RUN cargo add --quiet launchdarkly-server-sdk \
    && cargo add --quiet tokio@1 -F rt,macros \
    && cargo add --quiet futures \
    && cargo add --quiet launchdarkly-sdk-transport -F hyper \
    && cargo build --quiet

WORKDIR /work

COPY shared /harness-shared
COPY languages/rust/harness /harness

ENTRYPOINT ["/harness/run.sh"]
