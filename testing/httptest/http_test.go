package httptest

import (
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestMakeHanlde(t *testing.T) {
// 	s := makeHanlde(t, nil)

// 	exitSignature := "a49d4586ff218ee472db91b89e6ee5ff0f0f881ecd42fe5a79420efd44371be8cce5be25641ebb223aa5e8580ceec0b913c2d7f4077b21ebd55287451d2d2384b7813382aa6860a3785fae9a23460a1bbfaf483f86db5ec68ac502f945a88423"

// 	_, err := crypto.EncryptUsingPublicKey([]byte(exitSignature), s.publickey)
// 	require.Nil(t, err)
// }

func TestReadFromFile(t *testing.T) {
	block, _ := pem.Decode([]byte(pemString))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	require.Nil(t, err)
	require.NotNil(t, key)
}
