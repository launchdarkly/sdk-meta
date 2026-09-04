ARG NODE_IMAGE
FROM ${NODE_IMAGE}

COPY shared/versions/npm.env /versions/npm.env
RUN . /versions/npm.env && npm install -g npm@"$NPM_VERSION" \
 && printf 'min-release-age=3\nmin-release-age-exclude[]=@launchdarkly/*\nmin-release-age-exclude[]=launchdarkly-*\nmin-release-age-exclude[]=highlight.run\n' > /root/.npmrc

WORKDIR /work

COPY shared /harness-shared
COPY languages/node/harness /harness

ENTRYPOINT ["/harness/run.sh"]
