package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type XMLSigned struct {
	XMLName   xml.Name  `xml:"root"`
	Signature Signature `xml:"Signature"`
	Challenge string    `xml:"challenge"`
}

type Signature struct {
	XMLName        xml.Name `xml:"Signature"`
	SignatureValue string   `xml:"SignatureValue"`
	KeyInfo        KeyInfo  `xml:"KeyInfo"`
}

type KeyInfo struct {
	XMLName  xml.Name `xml:"KeyInfo"`
	X509Data X509Data `xml:"X509Data"`
}

type X509Data struct {
	XMLName         xml.Name `xml:"X509Data"`
	X509Certificate string   `xml:"X509Certificate"`
}

func main() {
	// conf, err := settings.ParseYAML("config.yml")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// b, e := bridge.NewKalkanBridge()
	// if e != nil {
	// 	fmt.Println("here?", e)
	// 	return
	// }
	// defer b.Close()
	// b.Init()
	// b.KCLoadKeyStore(conf.DigitalSignaturePass, conf.DigitalSignaturePath)
	// // b.X509ExportCertificateFromStore()
	// s, rv := b.SignXML(`<company-id>770704034</company-id>`)
	// fmt.Println("SignXML", rv)
	// s, rv = b.VerifyXML(s)
	// fmt.Println("VerifyXML", rv)

	xmlFile, err := os.Open("cmd/cli/cert.xml")
	if err != nil {
		fmt.Println("Open", err)
		return
	}
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Readall", err)
		return
	}
	var xmlSigned XMLSigned
	err = xml.Unmarshal(byteValue, &xmlSigned)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(xmlSigned.Challenge)
	fmt.Println(xmlSigned.Signature.SignatureValue)
	fmt.Println(xmlSigned.Signature.KeyInfo.X509Data.X509Certificate)
	// fmt.Println("VerifyXML", )
	// fmt.Println("VerifyData", b.KC_GetLastErrorString())
}
