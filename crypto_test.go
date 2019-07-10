package main

import (
	"testing"
)

func TestHexToB64Encode(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	actual, _ := getB64FromHex(input)
	if actual != expected {
		t.Errorf("Base64 returned unexpected value. Expected %s, got %s.", expected, actual)
	}
	t.Logf("\ninput:    %s\nexpected: %s\nactual:   %s", input, expected, actual)
}

func TestXor(t *testing.T) {
	str1 := "1c0111001f010100061a024b53535009181c"
	str2 := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"
	actual, _ := xor(str1, str2)
	if actual != expected {
		t.Errorf("XOR returned unexpected value. Expected %s, got %s.", expected, actual)
	}
	t.Logf("\ninput:    %s,%s\nexpected: %s\nactual:   %s", str1, str2, expected, actual)
}
