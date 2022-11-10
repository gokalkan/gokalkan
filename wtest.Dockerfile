FROM ubuntu:20.04

# set timezone
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ="Asia/Almaty"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# install dependencies
ARG DEPS=" \
    git \
    gcc \
    libpcsclite-dev \
    zlib1g-dev \
    libltdl7 \
    curl \
    "
RUN apt-get update && apt-get install -y ${DEPS}

# install golang
ARG GO_VERSION=1.17.10

RUN curl -sL https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz -o /tmp/go${GO_VERSION}.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf /tmp/go${GO_VERSION}.linux-amd64.tar.gz && rm /tmp/go${GO_VERSION}.linux-amd64.tar.gz

# golang config
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# download go modules
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# install CA certs
ARG CA_DIRECTORY=./internal/ca-test

RUN mkdir /usr/local/share/ca-certificates/extra
COPY ${CA_DIRECTORY}/*.crt /usr/local/share/ca-certificates/extra/
COPY ${CA_DIRECTORY}/*.pem /etc/ssl/certs/
# RUN update-ca-certificates

# install kalkan libs
COPY ./sdk/libkalkancryptwr-64.so /usr/lib/

# build app
COPY . .
RUN go test -v -run . ./internal/testdata/wtest/
