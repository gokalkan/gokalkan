#ifndef KALKANCRYPT_H
#define KALKANCRYPT_H


#include <stdlib.h> 
#ifndef _WIN32
#include <time.h>
#endif
//---------------------------------------------------------------------------------------

#ifdef __cplusplus
extern "C" {
#endif

#ifdef _DLL
	#ifdef KALKAN_EXPORTS
		#define KC_DECL
	#else
		#define KC_DECL
	#endif
#else
#define KC_DECL
#endif

//---------------------------------------------------------------------------------------

/* Поддерживаемые носители:
	Пока работает только Файловая система (PKCS#12)
*/
	
#define KCST_PKCS12			0x00000001
#define KCST_KZIDCARD		0x00000002
#define KCST_KAZTOKEN		0x00000004
#define KCST_ETOKEN72K		0x00000008
#define KCST_JACARTA		0x00000010
#define KCST_X509CERT		0x00000020
#define KCST_AKEY			0x00000040
#define KCST_ETOKEN5110		0x00000080
	
//---------------------------------------------------------------------------------------
	/*
    Формат импорта/экспорта сертификата
	*/

#define KC_CERT_DER			0x00000101
#define KC_CERT_PEM			0x00000102
#define KC_CERT_B64			0x00000104

//---------------------------------------------------------------------------------------

#define KC_CERT_CA				0x00000201
#define KC_CERT_INTERMEDIATE	0x00000202
#define KC_CERT_USER			0x00000204

//---------------------------------------------------------------------------------------

#define KC_USE_NOTHING			0x00000401
#define KC_USE_CRL				0x00000402
#define KC_USE_OCSP				0x00000404

//---------------------------------------------------------------------------------------

#define KC_XML_INCL_C14N			0x01000001
#define KC_XML_INCL_C14NCOMMENT		0x01000002
#define KC_XML_INCL_C14N11			0x01000004
#define KC_XML_INCL_C14N11COMMENT	0x01000008
#define KC_XML_EXCL_C14N			0x01000010
#define KC_XML_EXCL_C14NCOMMENT		0x01000020


#define KC_XMLC_INCL_C14N			0x01000040
#define KC_XMLC_INCL_C14NCOMMENT	0x01000080
#define KC_XMLC_INCL_C14N11			0x01000100
#define KC_XMLC_INCL_C14N11COMMENT	0x01000200
#define KC_XMLC_EXCL_C14N			0x01000400
#define KC_XMLC_EXCL_C14NCOMMENT	0x01000800


//---------------------------------------------------------------------------------------

#define KC_CERTPROP_ISSUER_COUNTRYNAME		0x00000801
#define KC_CERTPROP_ISSUER_SOPN				0x00000802
#define KC_CERTPROP_ISSUER_LOCALITYNAME		0x00000803
#define KC_CERTPROP_ISSUER_ORG_NAME			0x00000804
#define KC_CERTPROP_ISSUER_ORGUNIT_NAME		0x00000805
#define KC_CERTPROP_ISSUER_COMMONNAME		0x00000806


#define KC_CERTPROP_SUBJECT_COUNTRYNAME		0x00000807
#define KC_CERTPROP_SUBJECT_SOPN			0x00000808
#define KC_CERTPROP_SUBJECT_LOCALITYNAME	0x00000809
#define KC_CERTPROP_SUBJECT_COMMONNAME		0x0000080a
#define KC_CERTPROP_SUBJECT_GIVENNAME		0x0000080b
#define KC_CERTPROP_SUBJECT_SURNAME			0x0000080c
#define KC_CERTPROP_SUBJECT_SERIALNUMBER	0x0000080d
#define KC_CERTPROP_SUBJECT_EMAIL			0x0000080e
#define KC_CERTPROP_SUBJECT_ORG_NAME		0x0000080f
#define KC_CERTPROP_SUBJECT_ORGUNIT_NAME	0x00000810
#define KC_CERTPROP_SUBJECT_BC				0x00000811
#define KC_CERTPROP_SUBJECT_DC				0x00000812
												   
												   
#define KC_CERTPROP_NOTBEFORE				0x00000813
#define KC_CERTPROP_NOTAFTER				0x00000814
												   
#define KC_CERTPROP_KEY_USAGE				0x00000815
#define KC_CERTPROP_EXT_KEY_USAGE			0x00000816
												   
#define KC_CERTPROP_AUTH_KEY_ID				0x00000817
#define KC_CERTPROP_SUBJ_KEY_ID				0x00000818
#define KC_CERTPROP_CERT_SN					0x00000819
											   
												   
#define KC_CERTPROP_ISSUER_DN				0x0000081a
#define KC_CERTPROP_SUBJECT_DN				0x0000081b
												   
#define KC_CERTPROP_SIGNATURE_ALG			0x0000081c

#define KC_CERTPROP_PUBKEY					0x0000081d

#define	KC_CERTPROP_POLICIES_ID				0x0000081e

//---------------------------------------------------------------------------------------

//--- KALKANCRYPTCOM_FLAGS ---
#define KC_SIGN_DRAFT		0x00000001
#define KC_SIGN_CMS			0x00000002
#define KC_IN_PEM			0x00000004
#define KC_IN_DER			0x00000008
#define KC_IN_BASE64		0x00000010
#define KC_IN2_BASE64		0x00000020
#define KC_DETACHED_DATA	0x00000040
#define KC_WITH_CERT		0x00000080
#define KC_WITH_TIMESTAMP	0x00000100
#define KC_OUT_PEM			0x00000200
#define KC_OUT_DER			0x00000400
#define KC_OUT_BASE64		0x00000800
#define KC_PROXY_OFF		0x00001000
#define KC_PROXY_ON			0x00002000
#define KC_PROXY_AUTH		0x00004000
#define KC_IN_FILE			0x00008000
#define KC_NOCHECKCERTTIME	0x00010000
#define KC_HASH_SHA256		0x00020000
#define KC_HASH_GOST95		0x00040000
#define KC_GET_OCSP_RESPONSE		0x00080000


//---------------------------------------------------------------------------------------

#define KCR_BASE				0x08F00000

#define KCR_OK							   0x00000000
#define KCR_INIT_ERROR			KCR_BASE + 0x00000001
#define KCR_ERROR_READ_PKCS12	KCR_BASE + 0x00000002
#define KCR_ERROR_OPEN_PKCS12	KCR_BASE + 0x00000003
#define KCR_INVALID_PROPID		KCR_BASE + 0x00000004
#define KCR_BUFFER_TOO_SMALL	KCR_BASE + 0x00000005
#define KCR_CERT_PARSE_ERROR	KCR_BASE + 0x00000006
#define KCR_INVALID_FLAG		KCR_BASE + 0x00000007
#define KCR_OPENFILEERR			KCR_BASE + 0x00000008
#define KCR_INVALIDPASSWORD		KCR_BASE + 0x00000009
#define KCR_CERTWRONGDATE		KCR_BASE + 0x0000000a
#define KCR_CERTEXPIRED			KCR_BASE + 0x0000000b
#define KCR_ISNOTCACERT			KCR_BASE + 0x0000000c
#define KCR_MEMORY_ERROR		KCR_BASE + 0x0000000d
#define KCR_CHECKCHAINERROR		KCR_BASE + 0x0000000e
#define KCR_CACERTKEYUSAGEERROR	KCR_BASE + 0x0000000f
#define KCR_VALIDTYPEERROR		KCR_BASE + 0x00000010
#define KCR_BADCRLFORMAT		KCR_BASE + 0x00000011
#define KCR_LOADCRLERROR		KCR_BASE + 0x00000012
#define KCR_LOADCRLSERROR		KCR_BASE + 0x00000013

#define KCR_UNKNOWN_ALG			KCR_BASE + 0x00000015
#define KCR_KEYNOTFOUND			KCR_BASE + 0x00000016
#define KCR_SIGN_INIT_ERROR		KCR_BASE + 0x00000017
#define KCR_SIGN_ERROR			KCR_BASE + 0x00000018
#define KCR_ENCODE_ERROR		KCR_BASE + 0x00000019
#define KCR_INVALID_FLAGS		KCR_BASE + 0x0000001a
#define KCR_CERTNOTFOUND		KCR_BASE + 0x0000001b
#define KCR_VERIFYSIGNERROR		KCR_BASE + 0x0000001c
#define KCR_BASE64_DECODE_ERROR	KCR_BASE + 0x0000001d
#define KCR_UNKNOWN_CMS_FORMAT	KCR_BASE + 0x0000001e
#define KCR_GETHASHERROR		KCR_BASE + 0x0000001f
#define KCR_CA_CERT_NOT_FOUND	KCR_BASE + 0x00000020
#define KCR_XMLSECINIT_ERROR	KCR_BASE + 0x00000021
#define KCR_LOADTRUSTEDCERTSERR	KCR_BASE + 0x00000022
#define KCR_SIGN_INVALID		KCR_BASE + 0x00000023
#define KCR_NOSIGNFOUND			KCR_BASE + 0x00000024
#define KCR_DECODE_ERROR		KCR_BASE + 0x00000025
#define KCR_XMLPARSEERROR		KCR_BASE + 0x00000026
#define KCR_XMLADDIDERROR		KCR_BASE + 0x00000027
#define KCR_XMLINTERNALERROR	KCR_BASE + 0x00000028
#define KCR_XMLSETSIGNERROR		KCR_BASE + 0x00000029
#define KCR_OPENSSLERROR		KCR_BASE + 0x0000002a
#define KCR_ENGINE_INITERR		KCR_BASE + 0x0000002b
#define KCR_NOTOKENFOUND		KCR_BASE + 0x0000002c
#define KCR_OCSP_ADDCERTERR		KCR_BASE + 0x0000002d
#define KCR_OCSP_PARSEURLERR	KCR_BASE + 0x0000002e
#define KCR_OCSP_ADDHOSTERR		KCR_BASE + 0x0000002f
#define KCR_OCSP_REQERR			KCR_BASE + 0x00000030
#define KCR_OCSP_CONNECTIONERR	KCR_BASE + 0x00000031
#define KCR_VERIFY_NODATA		KCR_BASE + 0x00000032
#define KCR_IDATTR_NOTFOUND		KCR_BASE + 0x00000033
#define KCR_IDRANGE				KCR_BASE + 0x00000034
#define KCR_XMLKEYDUPERROR		KCR_BASE + 0x00000035
#define KCR_XMLKEYCREATEERROR	KCR_BASE + 0x00000036
#define KCR_READERNOTFOUND		KCR_BASE + 0x00000037
#define KCR_GETCERTPROPERR		KCR_BASE + 0x00000038
#define KCR_SIGNFORMMAT			KCR_BASE + 0x00000039
#define KCR_INDATAFORMAT		KCR_BASE + 0x0000003a
#define KCR_OUTDATAFORMAT		KCR_BASE + 0x0000003b
#define KCR_VERIFY_INIT_ERROR	KCR_BASE + 0x0000003c
#define KCR_VERIFY_ERROR		KCR_BASE + 0x0000003d
#define KCR_HASH_ERROR          KCR_BASE + 0x0000003e
#define KCR_SIGNHASH_ERROR		KCR_BASE + 0x0000003f
#define KCR_CACERTNOTFOUND      KCR_BASE + 0x00000040
#define KCR_CERTTIMEINVALID     KCR_BASE + 0x00000042
#define KCR_CONVERTERROR        KCR_BASE + 0x00000043
#define KCR_TSACREATEQUERY		KCR_BASE + 0x00000044
#define KCR_CREATEOBJ			KCR_BASE + 0x00000045
#define KCR_CREATENONCE			KCR_BASE + 0x00000046
#define KCR_HTTPERROR			KCR_BASE + 0x00000047
#define KCR_CADESBES_FAILED		KCR_BASE + 0x00000048
#define KCR_CADEST_FAILED		KCR_BASE + 0x00000049
#define KCR_NOTSATOKEN			KCR_BASE + 0x0000004a
#define KCR_INVALID_DIGEST_LEN	KCR_BASE + 0x0000004b
#define KCR_GENRANDERROR		KCR_BASE + 0x0000004c
#define KCR_SOAPNSERROR			KCR_BASE + 0x0000004d
#define KCR_GETPUBKEY			KCR_BASE + 0x0000004e
#define KCR_GETCERTINFO			KCR_BASE + 0x0000004f
#define KCR_FILEREADERROR		KCR_BASE + 0x00000050
#define KCR_CHECKERROR			KCR_BASE + 0x00000051
#define KCR_ZIPEXTRACTERR		KCR_BASE + 0x00000052
#define KCR_NOMANIFESTFILE		KCR_BASE + 0x00000053



#define KCR_LIBRARYNOTINITIALIZED	KCR_BASE + 0x00000101

#define KCR_ENGINELOADERR		KCR_BASE + 0x00000200

#define KCR_PARAM_ERROR			KCR_BASE + 0x00000300

#define KCR_CERT_STATUS_OK		KCR_BASE + 0x00000400
#define KCR_CERT_STATUS_REVOKED	KCR_BASE + 0x00000401
#define KCR_CERT_STATUS_UNKNOWN	KCR_BASE + 0x00000402

//---------------------------------------------------------------------------------------


typedef struct stKCFunctions {
	unsigned long(*KC_Init)(void);

	unsigned long(*KC_GetTokens)(unsigned long storage, char *tokens, unsigned long *tk_count);
	unsigned long(*KC_GetCertificatesList)(char *certificates, unsigned long *cert_count);
	
	unsigned long(*KC_LoadKeyStore)(int storage, char *password, int passLen, char *container, int containerLen, char *alias);
	
	unsigned long(*X509LoadCertificateFromFile)(char *certPath, int certType);
	unsigned long(*X509LoadCertificateFromBuffer)(unsigned char *inCert, int certLength, int flag);
	unsigned long(*X509ExportCertificateFromStore)(char *alias, int flag, char *outCert, int *outCertLength);
	unsigned long(*X509CertificateGetInfo)(char *inCert, int inCertLength, int propId, unsigned char *outData, int *outDataLength);
	unsigned long(*X509ValidateCertificate)(char *inCert, int inCertLength, int validType, char *validPath, long long checkTime, char *outInfo, int *outInfoLength, int flags, char* getResp, int *getRespLength);
	
	unsigned long(*HashData)(char *algorithm, int flags, char *inData, int inDataLength, unsigned char *outData, int *outDataLength);
	unsigned long(*SignHash)(char *alias, int flags, char *inHash, int inHashLength, unsigned char *outSign, int *outSignLength);

	unsigned long(*SignData)(char *alias, int flags, char *inData, int inDataLength, unsigned char *inSign, int inSignLen, unsigned char *outSign, int *outSignLength);
	unsigned long(*SignDataArchive)(char *alias, int flags, char *inData, int inDataLength, unsigned char *inSign, int inSignLen, char *validPath, unsigned char *outSign, int *outSignLength);
	unsigned long(*SignXML)(char *alias, int flags, char *inData, int inDataLength, unsigned char *outSign, int *outSignLength, char *signNodeId, char *parentSignNode, char *parentNameSpace);

	unsigned long(*VerifyData)(char *alias, int flags, char *inData, int inDataLength, unsigned char *inoutSign, int inoutSignLength, char *outData, int *outDataLen, char *outVerifyInfo, int *outVerifyInfoLen, int inCertID, char *outCert, int *outCertLength);
	unsigned long(*VerifyXML)(char *alias, int flags, char *inData, int inDataLength, char *outVerifyInfo, int *outVerifyInfoLen);

	unsigned long(*KC_getCertFromXML)(const char* inXML, int inXMLLength, int inSignID, char *outCert, int *outCertLength);
	unsigned long(*KC_getSigAlgFromXML)(const char* xml_in, int xml_in_size, char *retSigAlg, int *retLen);

	unsigned long(*KC_GetLastError)(void);
	unsigned long(*KC_GetLastErrorString)(char *errorString, int *bufSize);
	void (*KC_XMLFinalize)(void);
	void (*KC_Finalize)(void);

	void(*KC_TSASetUrl)(char *tsaurl);
	unsigned long(*KC_GetTimeFromSig)(char *inData, int inDataLength, int flags, int inSigId, time_t *outDateTime);
	unsigned long (*KC_SetProxy)(int flags, char *inProxyAddr, char *inProxyPort, char *inUser, char *inPass);

	unsigned long (*KC_GetCertFromCMS)(char *inCMS, int inCMSLen, int inSignId, int flags, char *outCert, int *outCertLength);

	unsigned long(*SignWSSE)(char *alias, unsigned long flags, char *inData, int inDataLength, unsigned char *outSign, int *outSignLength, char *signNodeId);

	unsigned long (*ZipConVerify)(char *inZipFile, int flags, char *outVerifyInfo, int *outVerifyInfoLen);
	unsigned long (*ZipConSign)(char *alias, const char *filePath, const char *name, const char *outDir, int flags);
	unsigned long(*KC_getCertFromZipFile)(char* inZipFile, int flags, int inSignID, char *outCert, int *outCertLength);
	unsigned long (*UVerifyData)(char *alias, int flags, char *inData, int inDataLength, unsigned char *inOutSign, int inOutSignLength, char *outData, int *outDataLen, char *outVerifyInfo, int *outVerifyInfoLen, int inCertID, char *outCert, int *outCertLength);
	void(*KC_InitDebug)(void);

} stKCFunctionsType;

//---------------------------------------------------------------------------------------

//int KC_DECL dllmain(void);
int KC_DECL KC_GetFunctionList(stKCFunctionsType **KCfunc);

__attribute__((weak))
stKCFunctionsType *kc_funcs;

//---------------------------------------------------------------------------------------

#ifdef __cplusplus
}
#endif

#endif
