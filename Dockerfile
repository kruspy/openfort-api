FROM golang:1.19.1-buster AS builder

WORKDIR /opt

COPY cmd/openfort-api cmd/openfort-api
COPY api api
COPY pkg pkg
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile

RUN make build-linux

FROM alpine:3.16

WORKDIR /opt

RUN addgroup -g 1234 apigroup
RUN adduser -D apiuser -u 1234 -G apigroup

USER apiuser

COPY --from=builder /opt/out/bin/openfort-api openfort-api

CMD [ "./openfort-api" ]
