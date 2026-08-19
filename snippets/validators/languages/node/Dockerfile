FROM node:22

RUN npm install -g npm@12 \
 && printf 'min-release-age=3\nmin-release-age-exclude[]=@launchdarkly/*\nmin-release-age-exclude[]=launchdarkly-*\n' > /root/.npmrc

WORKDIR /work

COPY shared /harness-shared
COPY languages/node/harness /harness

ENTRYPOINT ["/harness/run.sh"]
