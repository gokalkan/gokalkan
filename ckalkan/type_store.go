package ckalkan

// StoreType представляет собой вид хранилища/носителя.
type StoreType int

const (
	StoreTypePKCS12     StoreType = 1   // Файловая система
	StoreTypeKZIDCard   StoreType = 2   // Удостоверение личности гражданина РК
	StoreTypeKazToken   StoreType = 4   // Казтокен
	StoreTypeEToken     StoreType = 8   // eToken 72k
	StoreTypeJaCarta    StoreType = 16  // JaCarta
	StoreTypeX509Cert   StoreType = 32  // Сертификат X509
	StoreTypeAKey       StoreType = 64  // aKey
	StoreTypeEToken5110 StoreType = 128 // eToken 5110
)
