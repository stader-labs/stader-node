package httptest

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"testing"

	_ "embed"

	"github.com/google/uuid"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	stadertypes "github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/stretchr/testify/require"
	types "github.com/wealdtech/go-eth2-types/v2"
	eth2ks "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
)

func TestMakeHanlde(t *testing.T) {
	s := makeHanlde(t, nil)

	exitSignature := "a49d4586ff218ee472db91b89e6ee5ff0f0f881ecd42fe5a79420efd44371be8cce5be25641ebb223aa5e8580ceec0b913c2d7f4077b21ebd55287451d2d2384b7813382aa6860a3785fae9a23460a1bbfaf483f86db5ec68ac502f945a88423"

	_, err := crypto.EncryptUsingPublicKey([]byte(exitSignature), s.publickey)
	require.Nil(t, err)
}

func TestReadFromFile(t *testing.T) {
	block, rest := pem.Decode([]byte(pemString))
	require.NotNil(t, block)
	require.Empty(t, rest)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	require.Nil(t, err)
	require.NotNil(t, key)
}

//go:embed validator.json
var validator []byte

var prv []byte = []byte{}

func TestValidatorKey(t *testing.T) {
	encryptor := eth2ks.New(eth2ks.WithCipher("scrypt"))

	var v validatorKey
	err := json.Unmarshal(validator, &v)
	require.Nil(t, err)

	b, err := encryptor.Decrypt(v.Crypto, "passhere")
	require.Nil(t, err)
	require.NotNil(t, b)

	prv, err := types.BLSPrivateKeyFromBytes(prv)
	fmt.Printf("PUB: %+v", stadertypes.BytesToValidatorPubkey(prv.PublicKey().Marshal()).String())
	fmt.Printf("prv: %#v", prv.Marshal())
	require.Nil(t, err)
	require.NotNil(t, prv)

}

type validatorKey struct {
	Crypto  map[string]interface{}      `json:"crypto"`
	Version uint                        `json:"version"`
	UUID    uuid.UUID                   `json:"uuid"`
	Path    string                      `json:"path"`
	Pubkey  stadertypes.ValidatorPubkey `json:"pubkey"`
}
