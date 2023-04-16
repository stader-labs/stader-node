package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func BytesToPublicKey(pub []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub)
	b := block.Bytes
	var err error

	key, err := x509.ParsePKCS1PublicKey(b)
	if err != nil {
		fmt.Printf("Error using x509.ParsePKIXPublicKey %v\n", err)
		return nil, err
	}

	return key, nil
}

func BytesToPrivateKey(pub []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pub)
	b := block.Bytes
	var err error

	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		fmt.Printf("Error using x509.ParsePKIXPublicKey %v\n", err)
		return nil, err
	}

	return key, nil
}

func EncryptUsingPublicKey(data []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	exitMsgEncrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, data, nil)
	if err != nil {
		return nil, err
	}

	return exitMsgEncrypted, nil
}

func DecryptUsingPublicKey(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	exitMsgEncrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, data, nil)
	if err != nil {
		return nil, err
	}

	return exitMsgEncrypted, nil
}
