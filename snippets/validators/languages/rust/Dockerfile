FROM rust:1.85

WORKDIR /work

COPY shared /harness-shared
COPY languages/rust/harness /harness

ENTRYPOINT ["/harness/run.sh"]
