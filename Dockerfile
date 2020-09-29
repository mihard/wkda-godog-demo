# syntax=docker/dockerfile:experimental
FROM golang:1.14 as build_env

ENV GO111MODULE=on
ENV GOPROXY=direct
ENV GOSUMDB=off

RUN apt-get update && apt-get install -y zip tree openssh-client ca-certificates

RUN go get github.com/cucumber/godog/cmd/godog@v0.10.0
RUN mkdir -m 0777 -p /go/src/github.com/mihard/wkda-godog-demo && chmod -R 0777 /go
WORKDIR /go/src/github.com/mihard/wkda-godog-demo

FROM build_env as build

RUN mkdir -m 700 /root/.ssh && \
    echo "Host *\n\tStrictHostKeyChecking no" >> /root/.ssh/config && \
    chmod 400 /root/.ssh/config && \
    echo "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig

WORKDIR /go/src/github.com/mihard/wkda-godog-demo
COPY . .
RUN --mount=type=ssh tree && go build -mod vendor -o /demo main.go

FROM debian:10-slim as service
COPY --from=build /demo /usr/bin/demo

CMD '/usr/bin/demo'

