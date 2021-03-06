#
# Copyright (C) 2018 Dell Technologies
#
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.11.2-alpine3.7 AS builder

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

RUN echo http://nl.alpinelinux.org/alpine/v3.7/main > /etc/apk/repositories; \
    echo http://nl.alpinelinux.org/alpine/v3.7/community >> /etc/apk/repositories

RUN apk update && apk add make && apk add bash
RUN apk add curl && apk add git && apk add openssh && apk add build-base

ADD getglide.sh .
RUN sh ./getglide.sh
# set the working directory
WORKDIR $GOPATH/src/github.com/edgexfoundry/device-snmp-patlite

COPY . .

RUN make prepare
RUN make build

FROM scratch

ENV APP_PORT=49992
EXPOSE $APP_PORT

COPY --from=builder /go/src/github.com/edgexfoundry/device-snmp-patlite/cmd /

ENTRYPOINT ["/device-snmp-patlite","--registry","--profile=docker","--confdir=/res"]

