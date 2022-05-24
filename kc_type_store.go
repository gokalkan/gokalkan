package gokalkan

type KCStoreType int // вид хранилища/носителя.

const (
	KCStoreTypePKCS12     KCStoreType = 1   // Файловая система
	KCStoreTypeKZIDCard   KCStoreType = 2   // Удостоверение личности гражданина РК
	KCStoreTypeKazToken   KCStoreType = 4   // Казтокен
	KCStoreTypeEToken     KCStoreType = 8   // eToken 72k
	KCStoreTypeJaCarta    KCStoreType = 16  // JaCarta
	KCStoreTypeX509Cert   KCStoreType = 32  // Сертификат X509
	KCStoreTypeAKey       KCStoreType = 64  // aKey
	KCStoreTypeEToken5110 KCStoreType = 128 // eToken 5110
)
