# syntax=docker/dockerfile:experimental
FROM golang:1.14 as build_env

ENV GO111MODULE=on
ENV GOPROXY=direct
ENV GOSUMDB=off

RUN apt-get update && apt-get install -y zip tree openssh-client

RUN go get github.com/cucumber/godog/cmd/godog@v0.10.0

FROM build_env as build

RUN mkdir -m 700 /root/.ssh && \
    echo "Host *\n\tStrictHostKeyChecking no" >> /root/.ssh/config && \
    chmod 400 /root/.ssh/config && \
    echo "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig

RUN mkdir -m 0777 -p /go/src/github.com/wkda/${GIT_REPO_NAME} \
    && chmod -R 0777 /go
WORKDIR /go/src/github.com/wkda/${GIT_REPO_NAME}
ADD . .
RUN chmod -R 0777 .

RUN --mount=type=ssh go build -o /demo main.go

FROM alpine:latest as service

COPY --from=build /demo /usr/bin/demo

CMD ['demo']

