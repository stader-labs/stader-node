package crypto

import "encoding/base64"

func DecodeBase64(data string) ([]byte, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return decodedPublicKey, nil
}

func EncodeBase64(data []byte) []byte {
	var dst []byte

	base64.StdEncoding.Encode(dst, data)

	return dst
}