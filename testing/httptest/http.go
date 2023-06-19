package httptest

import (
	"encoding/json"
	"net/http"
	"testing"

	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stretchr/testify/require"
)

type StaderHandler struct {
	data map[string]interface{}
	t    *testing.T
}

func SererHttp(t *testing.T) {
	mux := http.NewServeMux()
	s := StaderHandler{
		data: make(map[string]interface{}),
		t:    t,
	}
	mux.HandleFunc("/presigns", s.presigns)
	mux.HandleFunc("/publicKey", s.publicKey)

	err := http.ListenAndServe(":9989", mux)

	require.Nil(t, err)
}

func (s *StaderHandler) presigns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (s *StaderHandler) publicKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := stader_backend.PublicKeyApiResponse{
		Value: "true",
	}
	json.NewEncoder(w).Encode(p)
}
