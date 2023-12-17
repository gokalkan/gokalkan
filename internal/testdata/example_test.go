package testdata

import (
	"encoding/base64"
	"fmt"

	"github.com/doodocs/doodocs/pkg/gokalkan"
	"github.com/doodocs/doodocs/pkg/gokalkan/types"
)

func ExampleClient_Sign() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	cli.LoadKeyStore(keyPath, keyPassword)

	signData, _ := cli.Sign([]byte("Hello World!"), false, false)

	signInBase64 := base64.StdEncoding.EncodeToString(signData)

	fmt.Printf("Подписанные данные:\n%s", signInBase64)
	// Output:
	// Подписанные данные:
	// ........
}

func ExampleClient_Verify() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	cli.LoadKeyStore(keyPath, keyPassword)

	signData, _ := cli.Sign([]byte("Hello World!"), false, false)

	ver, _ := cli.Verify(&types.VerifyInput{
		SignatureBytes:    signData,
		DataBytes:         []byte("Hello World!"),
		IsDetached:        false,
		MustCheckCertTime: false,
	})
	fmt.Println(ver)
	// Output:
	//Signature N 1
	//Id = 1
	//certificateSerialNumber=.....................................
	//signatureAlgorithm=sha256WithRSAEncryption(1.2.840.113549.1.1.11)
	//serialNumber=IIN1234567891011
	//Signature is OK
}

func ExampleClient_GetTimeFromSig() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	cli.LoadKeyStore(keyPath, keyPassword)

	signData, _ := cli.Sign([]byte("Hello World!"), false, true)

	signInBase64 := base64.StdEncoding.EncodeToString(signData)

	time, _ := cli.GetTimeFromSig(signInBase64, true)

	fmt.Printf("Время подписания: %s\n", time)
	// Output:
	// Время подписания: time....
}

func ExampleClient_SignWSSE() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	// Подгружаем хранилище с паролем
	cli.LoadKeyStore(keyPath, keyPassword)

	signedWSSE, _ := cli.SignWSSE("<root>this is sample</root>", "1")

	fmt.Printf("Подписанный документ WSSE: %s\n", signedWSSE)
	// Output:
	// Подписанный документ WSSE:
}

func ExampleClient_X509ExportCertificateFromStore() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	// Подгружаем хранилище с паролем
	cli.LoadKeyStore(keyPath, keyPassword)

	cert, _ := cli.X509ExportCertificateFromStore(true)

	fmt.Printf("Сертификат: %s\n", cert)
	// Output:
	// Сертификат:
}

func ExampleClient_X509CertificateGetInfo() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	// Подгружаем хранилище с паролем
	cli.LoadKeyStore(keyPath, keyPassword)

	// Заполняем необходимые поля
	fields := []string{
		"CertPropIssuerCountryName",
		"CertPropNotAfter",
		"CertPropOCSP",
		"CertPropSubjectSurname",
	}

	// Экспортируем сертификат из хранилища
	cert, _ := cli.X509ExportCertificateFromStore(true)

	info, _ := cli.X509CertificateGetInfo(cert, fields)

	fmt.Printf("Информация по сертификату: %s\n", info)
	// Output:
	// Информация по сертификату:
}

func ExampleClient_GetCertFromXML() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	// Подгружаем хранилище с паролем
	cli.LoadKeyStore(keyPath, keyPassword)

	signedXML, _ := cli.SignXML("<root>this is a sample</root>", true)

	cert, _ := cli.GetCertFromXML(signedXML, 0)

	fmt.Printf("Сертификат: %s\n", cert)
	// Output:
	// Информация по сертификату:
}
