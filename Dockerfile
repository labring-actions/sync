FROM alpine:3.17

ARG TARGETARCH
WORKDIR /

COPY bin/sync-$TARGETARCH /usr/bin/sync
COPY bin/sealos /usr/bin

ENTRYPOINT ["/usr/bin/sync"]