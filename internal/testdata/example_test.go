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
