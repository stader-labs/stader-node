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

	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		fmt.Printf("Error using x509.ParsePKIXPublicKey %v\n", err)
		return nil, err
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		fmt.Printf("Error validating public key %v\n", err)
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
