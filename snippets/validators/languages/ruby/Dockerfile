ARG RUBY_IMAGE
FROM ${RUBY_IMAGE}

WORKDIR /work

COPY shared /harness-shared
COPY languages/ruby/harness /harness

ENTRYPOINT ["/harness/run.sh"]
