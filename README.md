# gokalkan

[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]

<img src="https://user-images.githubusercontent.com/29381624/170255957-56b2c349-c04f-4ec8-9054-78bbe351fcc8.png" width=100>

Пакет `gokalkan` является оберткой над нативными методами библиотеки KalkanCrypt.

KalkanCrypt является основной библиотекой для работы с ЭЦП ключами от pki.gov.kz и
позволяет подписывать, верифицировать, вытаскивать данные и много чего другого связанного
с электронными подписями, цифровыми сертификатами и ключами.

Особенности пакета `gokalkan`:

- Поддержка мультипоточности
- Без зависимостей
- Чистый код
- Напрямую вызывает нативные методы KalkanCrypt

## Перед использованием

Чтобы использовать библиотеку требуется провести подготовку:

#### 1. Получить SDK

Обратиться в [pki.gov.kz](https://pki.gov.kz/developers/) чтобы получить SDK. SDK представляет из себя набор библиотек для Java и C.

#### 2. Установить CA сертификаты

Сертификаты лежат по пути `SDK/C/Linux/ca-certs/Ubuntu`. В папке находятся два типа сертификатов - `production` и `test`. Для их установки приготовлены скрипты в той же директории. При запуске понадобятся sudo права.

#### 3. Скопировать .so файлы

Файлы лежат в директории `SDK/C/Linux/C`. Команда для копирования:

```sh
sudo cp -f libkalkancryptwr-64.so libkalkancryptwr-64.so.1.1.0 /usr/lib/
```

#### 4. Скопировать kalkancrypt

Скопировать папку `SDK/C/Linux/libs_for_linux/kalkancrypt` в `/opt/`:

```sh
sudo cp -r kalkancrypt /opt/
```

#### 5. Настроить права доступа kalkancrypt

```sh
sudo chmod -R 555 /opt/kalkancrypt
```

#### 6. Установить переменную окружения

При использовании `gokalkan` убедитесь, что экспортирована переменная окружения:

```sh
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/opt/kalkancrypt/:/opt/kalkancrypt/lib/engines
```

Эта переменная нужна для обращения к библиотеке KalkanCrypt.

## Установка

Версия Go 1.17+

```sh
go get github.com/gokalkan/gokalkan
```

## Загрузка хранилища PKCS12

Загрузка хранилища с ключом и сертификатом (например ЭЦП ключ, который начинается с `RSA...`):

```go
package main

import (
	"fmt"
	"log"

	"github.com/gokalkan/gokalkan"
)

var (
	certPath = "test_cert/GOSTKNCA.p12" // путь к хранилищу

	certPassword = "Qwerty12" // пароль
	// P.S. никогда не храните пароли в коде
)

func main() {
	// для теста
	opts := gokalkan.OptsTest

	// для прода
	// opts := gokalkan.OptsProd

	cli, err = gokalkan.NewClient(opts...)
	if err != nil {
		log.Fatal("new kalkan client create error", err)
	}
	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	// Подгружаем хранилище с паролем
	err = cli.LoadKeyStore(certPath, certPassword)
	if err != nil {
		log.Fatal("load key store error", err)
	}
}
```

Следует отметить, что при инициализации gokalkan клиента нужно указывать опцию.
Есть две опции - `OptsTest` и `OptsProd`. Нужно выбрать одну из них в зависимости от окружения.

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

### Подпись XML документа для SmartBridge

Для того чтобы подписать XML документ в формате SignWSSE, нужно передать документ в виде строки.
Функция обернет документ в `soap:Envelope` и запишет внутри `soap:Body`.

```go
signedXML, err := cli.SignWSSE("<root>GoKalkan</root>")

fmt.Println("Подписанный XML в формате WSSE", signedXML)
fmt.Println("Ошибка", err)
```

## Бенчмарки

Команда запуска бенчмарка:

```sh
go test -bench SignXML -run=^$ -benchmem
```

Характеристики хост машины:

- goos: linux
- goarch: amd64
- cpu: Intel(R) Core(TM) i5-8500 CPU @ 3.00GHz

| Бенчмарк           | Кол-во циклов | Средн. время выполнения | Средн. потребление ОЗУ | Средн. кол-во аллокаций |
| ------------------ | ------------- | ----------------------- | ---------------------- | ----------------------- |
| BenchmarkSignXML-6 | 2809          | 422310 ns/op            | 2792 B/op              | 8 allocs/op             |

## Contributors ✨

Cпасибо за помощь в развитии проекта:

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<table>
	<tr>
		<td align="center">
			<a href="https://github.com/atlekbai">
				<img src="https://avatars.githubusercontent.com/u/29381624?v=4&s=100" width="100px;" alt=""/><br />
				<sub><b>Tlekbai Ali</b></sub>
			</a><br />
			<a href="https://github.com/gokalkan/gokalkan/commits?author=atlekbai" title="Code">💻</a>
		</td>
		<td align="center">
			<a href="https://github.com/gammban">
				<img src="https://avatars.githubusercontent.com/u/98373125?v=4s=100" width="100px;" alt=""/><br />
				<sub><b>Kilibayev Azat</b></sub>
			</a><br />
			<a href="https://github.com/gokalkan/gokalkan/commits?author=gammban" title="Code">💻</a>
		</td>
	</tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

## License

The MIT License (MIT) 2021 - [Abylaikhan Zulbukharov](https://github.com/Zulbukharov).

Please have a look at the [LICENSE.md](https://github.com/Zulbukharov/kalkancrypt-wrapper/blob/master/LICENSE.md) for more details.

[pkg-img]: https://pkg.go.dev/badge/Zulbukharov/GoKalkan
[pkg-url]: https://pkg.go.dev/github.com/gokalkan/gokalkan
[reportcard-img]: https://goreportcard.com/badge/Zulbukharov/GoKalkan
[reportcard-url]: https://goreportcard.com/report/Zulbukharov/GoKalkan
