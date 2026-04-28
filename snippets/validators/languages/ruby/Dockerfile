FROM ruby:3.3

WORKDIR /work

COPY shared /harness-shared
COPY languages/ruby/harness /harness

ENTRYPOINT ["/harness/run.sh"]
