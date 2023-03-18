package validator

// ROCKETPOOL-OWNED

// BLS signing root with domain
type signingRoot struct {
	ObjectRoot []byte `ssz-size:"32"`
	Domain     []byte `ssz-size:"32"`
}
