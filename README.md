# Golang Kalkancrypt Wrapper

⭐ Star us on GitHub — it motivates us a lot!

## Overview

<img href="logo.png"/>

Golang Kalkancrypt Wrapper - это простой веб-сервис для аутентификации посредством взаимодейсвия с \
с библиотеками kalkancrypt, используя ЭЦП.

## Kalkancrypt

Kalkancrypt - это набор библиотек для шифрования, дешифрования данных.
Одна из библиотек калкан это `libkalkancryptwr-64` файл с доступными методами для подписания файлов, \ 
текста используя ЭЦП. Подробнее про PKI можно почитать [здесь](lib/README.md).

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

