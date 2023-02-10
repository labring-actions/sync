FROM ghcr.io/labring/sealos:v4.1.4
ARG TARGETARCH
WORKDIR /

COPY bin/sync-$TARGETARCH /usr/bin/sync

ENTRYPOINT ["/usr/bin/sync"]