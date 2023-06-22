package httptest

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"testing"

	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/stretchr/testify/require"
)

type StaderHandler struct {
	data   map[string]interface{}
	t      *testing.T
	pubPEM string
	keyPEM []byte
}

func makeHanlde(t *testing.T) StaderHandler {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.Nil(t, err)

	publickey := &privatekey.PublicKey

	// Encode private key to PKCS#1 ASN.1 PEM.
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privatekey),
		},
	)

	// Encode public key to PKCS#1 ASN.1 PEM.
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publickey),
		},
	)

	pubkey := crypto.EncodeBase64(pubPEM)

	s := StaderHandler{
		data:   make(map[string]interface{}),
		t:      t,
		pubPEM: pubkey,
		keyPEM: keyPEM,
	}

	return s
}

func SererHttp(t *testing.T) {
	mux := http.NewServeMux()
	s := makeHanlde(t)
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
	json.NewDecoder(r.Body).Decode(&validatorPubKeys)

	var p map[string]bool
	// for _, v := range validatorPubKeys {
	// 	p[v.String()] = false
	// }
	json.NewEncoder(w).Encode(p)
}

func (s *StaderHandler) msgSubmitted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var validatorPubKeys []types.ValidatorPubkey
	json.NewDecoder(r.Body).Decode(&validatorPubKeys)

	var p map[string]bool
	for _, v := range validatorPubKeys {
		p[v.String()] = false
	}
	json.NewEncoder(w).Encode(p)
}

func (s *StaderHandler) presignsSubmitted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var validatorPubKeys stader_backend.BulkPreSignCheckApiRequestType
	json.NewDecoder(r.Body).Decode(&validatorPubKeys)

	p := make(map[string]bool)
	for _, v := range validatorPubKeys.ValidatorPubKeys {
		p[v.String()] = false
	}
	json.NewEncoder(w).Encode(p)
}

func (s *StaderHandler) presigns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (s *StaderHandler) presign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (s *StaderHandler) publicKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var validatorPubKeys map[string]interface{}
	json.NewDecoder(r.Body).Decode(&validatorPubKeys)

	p := stader_backend.PublicKeyApiResponse{
		Value: s.pubPEM,
	}
	json.NewEncoder(w).Encode(p)
}
