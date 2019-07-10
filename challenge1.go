package main

import (
	"encoding/base64"
	"encoding/hex"
)

// CRYPTO PALS SET 1 CHALLENGE 1
//     "Convert Hex to Base64"
//     https://cryptopals.com/sets/1/challenges/1
//
// The Hex string to decode:
//    49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
// Should produce:
//    SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

func getB64FromHex(hexString string) (string, error) {
	src := []byte(hexString)
	dst := make([]byte, hex.DecodedLen(len(src)))
	if _, err := hex.Decode(dst, src); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(dst), nil
}
