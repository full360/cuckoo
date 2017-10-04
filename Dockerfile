FROM alpine:3.6

ENV CUCKOO_VERSION 0.3.1

RUN apk add --no-cache \
      ca-certificates

RUN apk add --no-cache --virtual .build-deps \
      curl \
      unzip \
  && curl -fSL https://github.com/full360/cuckoo/releases/download/v${CUCKOO_VERSION}/cuckoo_${CUCKOO_VERSION}_linux_amd64.zip \
    -o /usr/local/bin/cuckoo.zip \
  && cd /usr/local/bin \
  && unzip cuckoo.zip \
  && rm cuckoo.zip \
  && curl -fSL http://pki.full360.com/full360-root-ca.crt \
    -o /usr/local/share/ca-certificates/full360-root-ca.crt \
  && update-ca-certificates \
  && apk del .build-deps
