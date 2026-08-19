ARG PHP_IMAGE
FROM ${PHP_IMAGE}

RUN apt-get update && apt-get install -y --no-install-recommends \
    git unzip ca-certificates curl python3 \
 && rm -rf /var/lib/apt/lists/* \
 && curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

WORKDIR /work

COPY shared /harness-shared
COPY languages/php/harness /harness

ENTRYPOINT ["/harness/run.sh"]
