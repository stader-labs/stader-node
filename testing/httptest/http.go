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
	mux.HandleFunc("/presign", s.presign)
	mux.HandleFunc("/presignsSubmitted", s.presignsSubmitted)

	err := http.ListenAndServe(":9989", mux)

	require.Nil(t, err)
}

func (s *StaderHandler) presignsSubmitted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p map[string]bool
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
	p := stader_backend.PublicKeyApiResponse{
		Value: "LS0tLS1CRUdJTiBSU0EgUFVCTElDIEtFWS0tLS0tCk1JSUNDZ0tDQWdFQXJIUXhMQXVuY0J3blFmUlo0QTdneXlhTHZ5TjFjenMrbWJrTzdLZkx3L25SVXFITTZ6ak4Kd1pSbXM4SDFMQ0tCVnhUMWVBYVhWZ05RSlE2NmgvYnYxYS9GSHdzU1Z2eDRyWlZ1SXRTN0hVMWNFTm5YYXZpbwprcGYrbW9odXE1UC9zaDVRY3FKTjZJU3JhRmlyZHI1allNUXhIUE9SV0M5aEhZTjJxZE9RUmRCMmhxN0Y3UThwCmxIT0FDQlI4VTUvTm43WDMzWkpiWVdtekF2dE10K3JFMVBWWi9LczY5dm84Yyt4bTFQRFpsZVYxQTgzUGdEZDMKOURWMXZQSS9pZXE3OWRzQVQyQVZTVVpoVm5kSExhaXdvczZIcDNVUHlXbnpaQTVoNGRCblo5Nk53bGw4VCsrOApyZ3MrbldrQVU4dXg3dUhnWjlyeVJUR2NudnZubFdmUEpvUThDVDRCWVN6QURWMUphMnUvN3pXZ0VQZm1EK0kyCnQ5ejU1eGJpZis2QVhHYXY0dHRITUs3dzB1aU1nZnMrU1J6SU0xSDJXZklXeFVzVTBNQnFCYWIvMmdSb3k4dkMKWFE0ZVJFNWpvdkR5OVNpMUlzbzZYTTc5Vm1FWklseWFzS2NLUlRZMTRQd1dJVEFQek5xRThub01QLzJTcmhJNwpiWUhJY3d6Z0pGbk15b3g4QXM4T3p3SlMzY1lwZTk5cHVGblVoWmxrTjRIUjd5cEdOV0luRVZBcEl5aFVCMVd6CjlJcmpTNHR6THoyVUhZdTNKOFpMUmdlWERxSVdNbUlmWk5wazFETUV1eFo1eFFtUS9oR0toVVBlZGo5NVJ0a2cKUUtCeWN6QnRndHorSE11SUFJOHpFcUJ1WmNkWS9zakFnZjM5QVU1RXdhQ0JqSFJIS0NyR3A0RUNBd0VBQVE9PQotLS0tLUVORCBSU0EgUFVCTElDIEtFWS0tLS0t",
	}
	json.NewEncoder(w).Encode(p)
}
