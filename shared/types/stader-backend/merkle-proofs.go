package stader_backend

type CycleMerkleProofs struct {
	Root  string   `json:"root"`
	Eth   string   `json:"eth"`
	Sd    string   `json:"sd"`
	Proof []string `json:"proof"`
	Cycle int64    `json:"cycle"`
}
