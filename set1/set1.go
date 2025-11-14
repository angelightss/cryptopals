package set1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func HexToBase64(hx string) (string, error) {
	v, err := hex.DecodeString(hx)
	if err != nil {
		return "", err
	}
	log.Printf("%s", v)
	return base64.StdEncoding.EncodeToString(v), nil
}

func XOR(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("xor: mismatched lenghts")
	}
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res
}
