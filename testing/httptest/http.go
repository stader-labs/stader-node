package httptest

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/eth2"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
)

type StaderHandler struct {
	data       map[string]interface{}
	t          *testing.T
	pubPEM     []byte
	keyPEM     []byte
	publickey  *rsa.PublicKey
	privatekey *rsa.PrivateKey
	bc         *services.BeaconClientManager
}

var (
	UserSettingPath = filepath.Join(ConfigPath, "user-settings.yml")
	ConfigPath, _   = homedir.Expand("~/.stader_testing")
)

func (s *StaderHandler) signatureDomain(t *testing.T, exitEpoch uint64) []byte {

	signatureDomain, err := s.bc.GetDomainData(eth2types.DomainVoluntaryExit[:], exitEpoch, false)

	require.Nil(t, err)

	return signatureDomain
}

var (
	// filename  = "localocal-testnet-presign-public-key"
	bitSize   = 4096
	pemString = `-----BEGIN RSA PRIVATE KEY-----
	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
	-----END RSA PRIVATE KEY-----
	`
)

func makeHanlde(t *testing.T, bc *services.BeaconClientManager) StaderHandler {
	block, _ := pem.Decode([]byte(pemString))
	privatekey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	require.Nil(t, err)

	// Encode private key to PKCS#1 ASN.1 PEM.
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: block.Bytes,
		},
	)

	publickey := &privatekey.PublicKey

	pubkeyIX, err := x509.MarshalPKIXPublicKey(publickey)
	require.Nil(t, err)

	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkeyIX,
		},
	)

	s := StaderHandler{
		data:       make(map[string]interface{}),
		t:          t,
		pubPEM:     pubPEM,
		keyPEM:     keyPEM,
		publickey:  publickey,
		privatekey: privatekey,
		bc:         bc,
	}

	return s
}

func SererHttp(t *testing.T, bc *services.BeaconClientManager) {
	mux := http.NewServeMux()
	s := makeHanlde(t, bc)
	mux.HandleFunc("/presign", s.presign)
	mux.HandleFunc("/presigns", s.presigns)
	mux.HandleFunc("/msgSubmitted", s.msgSubmitted)
	mux.HandleFunc("/presignsSubmitted", s.presignsSubmitted)
	mux.HandleFunc("/publicKey", s.publicKey)
	mux.HandleFunc("/merklesForElRewards", s.merklesForElRewards)

	err := http.ListenAndServe(":9989", mux)

	require.Nil(t, err)
}

func (s *StaderHandler) merklesForElRewards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var validatorPubKeys map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&validatorPubKeys)
	require.Nil(s.t, err)

	var p map[string]bool
	// for _, v := range validatorPubKeys {
	// 	p[v.String()] = false
	// }
	json.NewEncoder(w).Encode(p)
}

func (s *StaderHandler) msgSubmitted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var validatorPubKeys []types.ValidatorPubkey
	err := json.NewDecoder(r.Body).Decode(&validatorPubKeys)
	require.Nil(s.t, err)

	var p map[string]bool
	for _, v := range validatorPubKeys {
		_, ok := s.data[v.String()]
		p[v.String()] = ok
	}
	json.NewEncoder(w).Encode(p)
}

func (s *StaderHandler) presignsSubmitted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var validatorPubKeys stader_backend.BulkPreSignCheckApiRequestType
	err := json.NewDecoder(r.Body).Decode(&validatorPubKeys)
	require.Nil(s.t, err)

	p := make(map[string]bool)
	for _, v := range validatorPubKeys.ValidatorPubKeys {
		_, ok := s.data[v.String()]
		p[v.String()] = ok
	}
	json.NewEncoder(w).Encode(p)
}

func (s *StaderHandler) presigns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var preSignedMessages []stader_backend.PreSignSendApiRequestType
	err := json.NewDecoder(r.Body).Decode(&preSignedMessages)
	require.Nil(s.t, err)

	preSignSendResponse := make(map[string]stader_backend.PreSignSendApiResponseType, len(preSignedMessages))
	for _, v := range preSignedMessages {
		decodeSig, err := crypto.DecodeBase64(v.Signature)
		require.Nil(s.t, err)

		byteSig, err := crypto.DecryptUsingPublicKey(decodeSig, s.privatekey)
		require.Nil(s.t, err)

		var sig bls.Sign
		err = sig.DeserializeHexStr(string(byteSig))
		require.Nil(s.t, err)

		rootHash := s.srHash(s.t, v)

		var pub bls.PublicKey
		err = pub.DeserializeHexStr(v.ValidatorPublicKey)
		require.Nil(s.t, err)

		require.Nil(s.t, err)
		verify := sig.VerifyHash(&pub, rootHash[:])
		assert.True(s.t, verify)

		fmt.Printf("Success verify signature with pubkey: [%+v] [%+v] [%+v]\n", pub, sig, rootHash)

		require.Nil(s.t, err)
		s.data[v.ValidatorPublicKey] = true
		preSignSendResponse[v.ValidatorPublicKey] = stader_backend.PreSignSendApiResponseType{
			Success: true,
			Error:   "",
		}
	}

	json.NewEncoder(w).Encode(preSignSendResponse)
}

func (s *StaderHandler) presign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (s *StaderHandler) publicKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := stader_backend.PublicKeyApiResponse{
		Value: crypto.EncodeBase64(s.pubPEM),
	}
	json.NewEncoder(w).Encode(p)
}

func (s *StaderHandler) srHash(t *testing.T, request stader_backend.PreSignSendApiRequestType) [32]byte {
	epoch, err := strconv.ParseUint(request.Message.Epoch, 10, 64)
	require.Nil(t, err)

	validatorIndex, err := strconv.ParseUint(request.Message.ValidatorIndex, 10, 64)
	require.Nil(t, err)

	exitMessage := eth2.VoluntaryExit{
		Epoch:          epoch,
		ValidatorIndex: validatorIndex,
	}

	// Get object root
	or, err := exitMessage.HashTreeRoot()
	require.Nil(t, err)

	// Get signing root
	sr := eth2.SigningRoot{
		ObjectRoot: or[:],
		Domain:     s.signatureDomain(t, epoch),
	}

	srHash, err := sr.HashTreeRoot()
	require.Nil(t, err)

	return srHash
}
