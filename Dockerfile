FROM alpine:3.17

ARG TARGETARCH
WORKDIR /

COPY bin/sync-$TARGETARCH /sync

RUN wget https://github.com/labring/sealos/releases/download/v4.1.4/sealos_4.1.4_linux_amd64.tar.gz -o sealos.tar.gz  \
    && tar zxvpf sealos.tar.gz \
    && rm sealos.tar.gz \
    && cp sealos /usr/bin \

ENTRYPOINT ["/sync"]