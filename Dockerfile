FROM alpine:3.17

ARG TARGETARCH
WORKDIR /

COPY bin/sync-$TARGETARCH /sync

RUN wget -O sealos.tar.gz https://github.com/labring/sealos/releases/download/v4.1.4/sealos_4.1.4_linux_amd64.tar.gz

RUN tar zxvpf sealos.tar.gz
RUN rm sealos.tar.gz
RUN cp sealos /usr/bin

ENTRYPOINT ["/sync"]