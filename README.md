# Golang Kalkan Bindings

## Overview

Библиотека калкан это so файл с доступными методами для подписания файлов, текста \
используя ЭЦП. Принципы работы в PKI доступны [здесь]().

## About

Golang Kalkan Bindings - это обертка для библиотека на Golang.

## Usage

> Для запуска программы, необходимо:

- cp -r kalkancrypt /opt/
- добавить в доверенные сертификаты из ca_certs, которые находятся в SDK
- поместить `libkalkancryptwr-64.so` в директорию ./lib
- export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/opt/kalkancrypt/:/opt/kalkancrypt/lib/engines
- запустить `go run cmd/cli/main.go`

## Example

```go
func main() {
    fmt.Println("Hello World")
}
```

## Operation principles

