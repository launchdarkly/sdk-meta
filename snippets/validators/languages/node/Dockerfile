FROM node:20

WORKDIR /work

COPY shared /harness-shared
COPY languages/node/harness /harness

ENTRYPOINT ["/harness/run.sh"]
