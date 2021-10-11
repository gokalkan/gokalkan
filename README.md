# Golang Kalkancrypt Wrapper

<img src="assets/logo.png" width="200px" align='right'/>

⭐ Star on GitHub — it motivates me a lot!

## Overview

Golang Kalkancrypt Wrapper - это простой веб-сервис для аутентификации посредством взаимодейсвия с \
с библиотеками kalkancrypt, используя ЭЦП.

## Kalkancrypt

Kalkancrypt - это набор библиотек для шифрования, дешифрования данных.
Одна из библиотек калкан это `libkalkancryptwr-64` файл с доступными методами для подписания файлов, \ 
текста используя ЭЦП. Подробнее про PKI можно почитать [здесь](lib/README.md).

## Features

- Подписания текста, получения ответа в виде xml.
- Проверка XML подписи.

## Usage

> Для запуска программы, необходимо:

- Скопировать файлы kalkancrypt. \
`bash scripts/copy_libs.sh`\
**hint**  для получения SDK нужно обратиться в [pki.gov.kz](https://pki.gov.kz/developers/)
- Добавить в доверенные сертификаты из certs, которые находятся в SDK [pki.gov.kz](https://pki.gov.kz/developers/) для получения.
*hint* `bash scripts/install_certs.sh`
- добавить переменную окружения `LD_LIBRARY_PATH` для доступа программе к SDK. \
`export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/opt/kalkancrypt/:/opt/kalkancrypt/lib/engines`
- заполнить файл `config.yml.example` и переименовать.
`mv config.yml.example config.yml`
- запустить `go run cmd/cli/main.go`

## License
