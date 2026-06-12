FROM rust:1.94

WORKDIR /work

COPY shared /harness-shared
COPY languages/rust/harness /harness

ENTRYPOINT ["/harness/run.sh"]
