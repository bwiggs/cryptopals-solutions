package main

import (
	"encoding/hex"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestTranspose(t *testing.T) {
	ciphertext := []byte("ABCABCABCABC")
	blocks := transpose(ciphertext, 3)
	if string(blocks[0]) != "AAAA" {
		t.Error("unexpected block")
	}
	if string(blocks[1]) != "BBBB" {
		t.Error("unexpected block")
	}
	if string(blocks[2]) != "CCCC" {
		t.Error("unexpected block")
	}
}

func TestRotKeyDecrypt(t *testing.T) {
	expectedText := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	ciphertext, _ := hex.DecodeString("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	key := []byte("ICE")
	plaintext := rotKeyDecrypt(ciphertext, key)
	if string(plaintext) != expectedText {
		t.Error("unexpected decryption")
	}
}

func TestHamming(t *testing.T) {
	if 37 != hamming([]byte("this is a test"), []byte("wokka wokka!!!")) {
		t.Error("unexpected hamming distance")
	}
}

func TestFindKeySize(t *testing.T) {
	ciphertext, _ := hex.DecodeString("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	if findKeySize(ciphertext) != 3 {
		t.Error("unexpected key size")
	}
}

func TestFindKey(t *testing.T) {
	expectedKey := byte(58)
	ciphertext, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if key := findKey(ciphertext); key != expectedKey {
		spew.Dump(key)
		t.Error("unexpected key " + string(key))
	}
}
