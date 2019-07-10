package main

import (
	"encoding/hex"
)

// CRYPTO PALS SET 1 CHALLENGE 2
//     "Fixed XOR"
//     https://cryptopals.com/sets/1/challenges/2
//
// Write a function that takes two equal-length buffers and produces their XOR combination.
//    String 1: 1c0111001f010100061a024b53535009181c
//    String 2: 686974207468652062756c6c277320657965
// Should produce:
//    746865206b696420646f6e277420706c6179

func xor(a, b string) (string, error) {
	aSrc := []byte(a)
	bSrc := []byte(b)
	aDst := make([]byte, hex.DecodedLen(len(aSrc)))
	bDst := make([]byte, hex.DecodedLen(len(bSrc)))
	if _, err := hex.Decode(aDst, aSrc); err != nil {
		return "", err
	}
	if _, err := hex.Decode(bDst, bSrc); err != nil {
		return "", err
	}
	r := make([]byte, len(aDst))
	for i := range aDst {
		r[i] = aDst[i] ^ bDst[i]
	}
	re := hex.EncodeToString([]byte(r))
	return re, nil
}
