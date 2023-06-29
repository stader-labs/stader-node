package httptest

import (
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
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

//go:embed localocal-testnet-presign-key.rsa
var pemString string

var (
	filename = "localocal-testnet-presign-key"
	bitSize  = 4096
)

func makeHanlde(t *testing.T, bc *services.BeaconClientManager) StaderHandler {
	block, _ := pem.Decode([]byte(pemString))

	privatekey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	require.Nil(t, err)

	b := x509.MarshalPKCS1PrivateKey(privatekey)
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: b,
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

func ServeHttpLocal(t *testing.T, bc *services.BeaconClientManager) {
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

		fmt.Printf("Success verify signature with pubkey: [%+v] hash: [%+v]\n", pub.GetHexString(), rootHash)

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