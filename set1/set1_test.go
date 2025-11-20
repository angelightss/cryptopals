package set1

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"log"
	"testing"
)

func TestChallenge1(t *testing.T) {

	r, err := HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		t.Error(err)
	}
	if r != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Error("wrong string", r)
	}
}

func TestChallenge2(t *testing.T) {
	res := XOR(hexDecode(t, "1c0111001f010100061a024b53535009181c"),
		hexDecode(t, "686974207468652062756c6c277320657965"))
	if !bytes.Equal(res, hexDecode(t, "746865206b696420646f6e277420706c6179")) {
		t.Errorf("wront result: %x", res)
	}
}

func hexDecode(t *testing.T, s string) []byte {
	v, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal("fail to decode hex:", s)
	}
	log.Print(v)
	return v
}

func corpusFromFile(t *testing.T, name string) map[rune]float64 {
	text, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatal("failed to open corpus file:", err)
	}
	return buildCorpus(string(text))
}

func TestChallenge3(t *testing.T) {
	c := corpusFromFile(t, "dracula.txt")
	for char, val := range c {
		t.Logf("%c: %.5f", char, val)
	}

	res := findSingleXORKey(hexDecode(t, "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"), c)
	t.Logf("%s", res)
}
