package challenge

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// Challenge ...
type Challenge interface {
	BuildChallenge() string
	GetSerial() string
	GetUUID() string
}

type xmlChallenge struct {
	Serial string
	UUID   string
	// Body   string
}

type cmsChallenge struct {
	Serial string
	UUID   string
}

// NewChallenge ...
func NewChallenge(serial string, challengeType string) Challenge {
	uuid := uuid.New().String()
	if challengeType == "cms" {
		return &cmsChallenge{serial, uuid}
	}
	return &xmlChallenge{serial, uuid}
}

func (x *xmlChallenge) BuildChallenge() string {
	return fmt.Sprintf("<root><challenge>%s</challenge></root>", x.UUID)
}

func (x *xmlChallenge) GetSerial() string { return x.Serial }
func (x *xmlChallenge) GetUUID() string   { return x.UUID }

func (c *cmsChallenge) BuildChallenge() string { return "" }

func (c *cmsChallenge) GetSerial() string { return c.Serial }
func (c *cmsChallenge) GetUUID() string   { return c.UUID }

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

// ValidateSign will validate xml sign by challenge uuid,
// signature. cert exist
// on success returns challenge uuid
func ValidateSign(xmlData []byte) (string, error) {
	var xmlSigned XMLSigned

	err := xml.Unmarshal(xmlData, &xmlSigned)
	if err != nil {
		return "", err
	}

	challenge := xmlSigned.Signature.SignatureValue
	if challenge == "" ||
		xmlSigned.Challenge == "" ||
		xmlSigned.Signature.KeyInfo.X509Data.X509Certificate == "" {
		err = errors.New("Invalid sign")
	}
	return challenge, err
}
