# Kalkan Description [WIP]

P12 file (auth_rsa) digital certificate (PKCS#12) contains:
- private and public key,
- information about the owner (name, email address, IIN, etc.)


```bash
openssl pkcs12 -info -in file.p12
keytool -list -v -keystore file.p12 -storepass
```

## References

- [Openssl CLI usage](https://gist.github.com/dreikanter/c7e85598664901afae03fedff308736b)
- [P12 file extension explained](https://findanyanswer.com/what-is-a-p12-file-used-for)
- [SSL cert explained](https://www.virtues.it/2015/07/ssl-certificates-explained/)
- [Common commands with openssl](https://www.sslshopper.com/article-most-common-openssl-commands.html)
- [PKCS12 obtain flow](http://itdoc.hitachi.co.jp/manuals/3021/30213B6100e/ITSD0055.HTM)
- [Demo NCAlayer](https://ncalayer-react.netlify.app/)
- [Digital signature realisation review](https://ct.kz/hall-of-fame-eds/vmpgovkz/)
- [Digital signature formats](https://habr.com/ru/company/aktiv-company/blog/191866/)
- [Sign with private key](https://gist.github.com/ezimuel/3cb601853db6ebc4ee49)

<!-- when info is used -->
```
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            ***
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = KZ, CN = ***(RSA)
        Validity
            Not Before: Jun 11 11:53:39 2021 GMT
            Not After : Jun 11 11:53:39 2022 GMT
        Subject: CN = ***, serialNumber = IIN***, C = KZ, GN = ***, emailAddress = ***@GMAIL.COM
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (2048 bit)
                Modulus:
                    ***
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment
            X509v3 Extended Key Usage:
                TLS Web Client Authentication
            X509v3 Authority Key Identifier:
                keyid:***

            X509v3 Subject Key Identifier:
                1***
            X509v3 Certificate Policies:
                Policy: 1.2.398.3.3.2.4
                  CPS: http://pki.gov.kz/cps
                  User Notice:
                    Explicit Text: http://pki.gov.kz/cps

            X509v3 CRL Distribution Points:

                Full Name:
                  URI:http://crl.pki.gov.kz/nca_rsa.crl
                  URI: http://crl1.pki.gov.kz/nca_rsa.crl

            X509v3 Freshest CRL:

                Full Name:
                  URI:http://crl.pki.gov.kz/nca_d_rsa.crl
                  URI:http://crl1.pki.gov.kz/nca_d_rsa.crl

            Authority Information Access:
                CA Issuers - URI:http://pki.gov.kz/cert/nca_rsa.cer
                OCSP - URI:http://ocsp.pki.gov.kz

    Signature Algorithm: sha256WithRSAEncryption
         ***
```

1. alice generates private key
2. alice generates CSR (certificate signing request)
3. alice sends CSR to CA
4. CA encrypts CSR with private key and creates CER
5. Bob receives CER from public repo
6. Bob checks validity of CER and encrypts with CER

https://www.virtues.it/2015/07/ssl-certificates-explained/


## Commands

- private key
- public key
- certificate = public key + data
- p12 = private key + certificate


### Get info about p12 file

```sh
openssl pkcs12 -info -in RSA256_113eeea2c5812cb1d63e93d6ab8b0b02805c7fbb.p12
```

### Extract private key from p12 file

```sh
openssl pkcs12 -in ../*.p12 -out privatekey.pem -nodes -nocerts
```

### Generate public key using private key

```sh
openssl rsa -in privatekey.pem -outform PEM -pubout -out public.pem
```

### Get public key from cert file

```sh
openssl x509 -pubkey -noout -in cert.cert  > pubkey.pem
```

### Read so file functions
```sh
readelf -Ws libkalkancryptwr-64.so
ldd /usr/bin/openssl
```
