FROM ubuntu:20.04
# set tz
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ="Asia/Almaty"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
# install kalkan deps
RUN apt-get update && apt-get install -y wget git gcc libpcsclite-dev zlib1g-dev libltdl7
# install golang
COPY internal/testdata/deploy/go1.17.9.linux-amd64.tar.gz /tmp/
RUN tar -C /usr/local -xzf /tmp/go1.17.9.linux-amd64.tar.gz && rm /tmp/go1.17.9.linux-amd64.tar.gz
# golang config
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
# download go modules
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
# install CA certs
RUN mkdir /usr/local/share/ca-certificates/extra
COPY ./internal/testdata/wtest/ca/*.crt /usr/local/share/ca-certificates/extra/
COPY ./internal/testdata/wtest/ca/*.pem /etc/ssl/certs/
RUN update-ca-certificates
# install kalkan libs
COPY internal/testdata/deploy/lib/libkalkancryptwr-64.so /usr/lib/
COPY internal/testdata/deploy/lib/libkalkancryptwr-64.so.1.1.1 /usr/lib/
COPY internal/testdata/deploy/lib/kalkancrypt /opt/kalkancrypt/
ENV LD_LIBRARY_PATH /opt/kalkancrypt:/opt/kalkancrypt/lib/engines:$LD_LIBRARY_PATH
# build app
COPY . .
RUN date
RUN cat /etc/timezone
RUN go test -v -run . ./internal/testdata/wtest/
