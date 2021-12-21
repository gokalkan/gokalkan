# GoKalkan

GoKalkan - это библиотека-обертка над KalkanCrypt для Golang.

KalkanCrypt - это набор библиотек для шифрования, дешифрования данных.

Основные методы KalkanCrypt реализованы в `libkalkancryptwr-64`. Это файл доступными методами 
для подписания файлов, текста используя ЭЦП. Подробнее про PKI можно почитать [здесь](wiki/README.md).

## Доступный функционал

```go
// Kalkan - интерфейс с методами KalkanCrypt
type Kalkan interface {
	Init() error
	LoadKeyStore(password, containerPath string) error
	SignXML(data string) (string, error)
	VerifyXML(xml string) (string, error)
	VerifyData(data string) (*VerifiedData, error)
	X509ExportCertificateFromStore() (string, error)
	GetLastErrorString() string
	Close() error
}
```

Не все доступные методы пока были реализованы. Для знакомства со всеми функциями перейти [сюда](cpp/KalkanCrypt.h).

## Запуск

Чтобы использовать библиотеку требуется провести подготовку:

#### 1. Обратиться в [pki.gov.kz](https://pki.gov.kz/developers/) чтобы получить SDK

SDK представляет собой набор библиотек для Java и C.

#### 2. Установить в доверенные сертификаты

Сертификаты будут лежать по пути `SDK/C/Linux/ca-certs/Ubuntu`. Будут два типа сертфикатов - `production` и `test`.

В папке будут скрипты для установки сертификатов, понадобится sudo права.

#### 3. Скопировать `libkalkancryptwr-64.so` и `libkalkancryptwr-64.so.1.1.0` в /usr/lib/

Файлы лежат в директории `SDK/C/Linux/C`. Команда для копирования:

```sh
sudo cp -f libkalkancryptwr-64.so libkalkancryptwr-64.so.1.1.0 /usr/lib/
```

#### 4. Скопировать `kalkancrypt`  в `/opt/`

`kalkancrypt` - представляет набор из общих библиотек и состоит из файлов расширения `.so`.

Скопируйте папку `SDK/C/Linux/libs_for_linux/kalkancrypt` в `/opt/`

```sh
sudo cp -r kalkancrypt /opt/
```

#### 5. LD_LIBRARY_PATH

При обращении к GoKalkan убедитесь что экспортирована переменная окружения

```sh
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/opt/kalkancrypt/:/opt/kalkancrypt/lib/engines
```

Это переменная нужна для динамического обращения к библиотеке KalkanCrypt.


## Примеры

Начнем с загрузки сертификатов (можно ЭЦП, который начинается с `RSA...`):

```go
package main

import (
	"fmt"
	"log"

	kalkan "github.com/Zulbukharov/GoKalkan"
)

var (
	// certPath хранит путь к сертификату 
	certPath = "test_cert/GOSTKNCA.p12"

	// certPassword пароль
	// P.S. никогда не храните пароли в коде
	certPassword = "Qwerty12"
)

func main() {
	cli, err := kalkan.NewClient()
	if err != nil {
		log.Fatal("NewClient", err)
	}
	// Обязательно закрывайте клиент, иначе приведет утечкам ресурсов
	defer cli.Close()

	// Подгружаем сертификат с паролем
	if err := cli.LoadKeyStore(certPassword, certPath); err != nil {
		log.Fatal("cli.LoadKeyStore", err)
	}
}
```

### Подпись XML документа

Для того чтобы подписать XML документ, нужно передать документ в виде строки:

```go
signedXML, err := cli.SignXML("<root>GoKalkan</root>")

fmt.Println("Подписанный XML", signedXML)
fmt.Println("Ошибка", err)
```

### Проверка подписи на XML документе

Проверка подписи документа вернет ошибку, если документ подписан неверно либо срок
у сертификата с которым подписан истёк.

```go
serial, err := cli.VerifyXML(signedXML)

fmt.Println("Серийный номер", serial)
fmt.Println("Ошибка", err)
```

## Для чего эта библиотека

GoKalkan можно использовать для:
- подписывания XML документов c помощью ЭЦП
- реализовывания авторизации через ЭЦП
- подпись документов для гос. сервисов

GoKalkan не является библиотекой для подписывания XML документов на SmartBridge.

## Особенности

Библиотека GoKalkan не работает мультипоточно, т.е. нельзя использовать один 
экземпляр между горутинами. Это объясняется тем, что код CGO работает в отдельном 
потоке thread.

Для того, чтобы использовать между горутинами, следует передавать по каналу заранее
запущенным горутинам, которые принимают запросы на выполнение операции с GoKalkan.

## License

The MIT License (MIT) 2021 - [Abylaikhan Zulbukharov](https://github.com/Zulbukharov).

Please have a look at the [LICENSE.md](https://github.com/Zulbukharov/kalkancrypt-wrapper/blob/master/LICENSE.md) for more details.

