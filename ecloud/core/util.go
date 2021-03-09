package core

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"

	json "github.com/json-iterator/go"
)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC/VpRysi0bPRLS7sbgQDJHo1MAt9/bK
+nwK5Pe3z0/O4cH5I/8kFNYy4yFsLMM+zyFvVw9C4wzjHaRcmEuF3ziJMC9PD5ufUWgfO
5nSGgZW1cmgjqnhcWJ3i+Azj72RnhKQRCn9DgJduEC9MiKfbyTICGd6FXf9cxb21nkxI7vtwIDAQAB
-----END PUBLIC KEY-----
`)

func RsaEncrypt(origData []byte) ([]byte, error) {
	//将密钥解析成公钥实例
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	//解析pem.Decode（）返回的Block指针实例
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)
	//RSA算法加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func Dump(a interface{}) string {
	if bytes, err := json.MarshalIndent(&a, "", "  "); err == nil {
		return fmt.Sprintf("\n%s", bytes)
	}
	return ""
}
