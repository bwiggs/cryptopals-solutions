package main

import (
	"fmt"
	"strings"
)

// transpose makes a block that is the first byte of every block, and a block that is the second byte of every block, and so on.
func transpose(buf []byte, blockSize int) [][]byte {
	b := [][]byte{}

	for i := 0; i < blockSize; i++ {
		b = append(b, []byte{})
	}

	for i := 0; i < len(buf); i++ {
		j := i % blockSize
		b[j] = append(b[j], buf[i])
	}
	return b
}

func findKey(ct []byte) byte {
	commonChars := strings.Split(strings.ToLower("ETAOIN SHRDLU"), "")

	var best byte
	var bestScore int

	for i := 0; i < 255; i++ {
		b := byte(i)
		buf := []byte{}
		for j := range ct {
			buf = append(buf, ct[j]^b)
		}
		var score int
		for k := range commonChars {
			score += strings.Count(strings.ToLower(string(buf)), commonChars[k])
		}
		if score > bestScore {
			best = b
			bestScore = score
		}
	}

	return best
}

func hamming(s1, s2 []byte) int {
	var b uint
	var dist int
	for i := range s1 {
		for b = 0; b < 8; b++ {
			shift := uint(1 << b)
			if uint(s1[i])&shift != uint(s2[i])&shift {
				dist++
			}
		}
	}
	return dist
}

func findKeySize(buf []byte) int {
	bestScore := 9.0
	var keySize int
	for i, ks := 0, 2; ks <= 40; i, ks = i+1, ks+1 {
		if len(buf) < 2*ks {
			break
		}
		chunk1, chunk2 := buf[0:ks], buf[ks:2*ks]
		dist := hamming(chunk1, chunk2)
		score := float64(dist) / float64(ks)
		fmt.Printf("ks: %02d\tdist: %02d\tscore: %f\n", ks, dist, score)
		if score < bestScore {
			bestScore = score
			keySize = ks
		}
	}
	return keySize
}

func rotKeyDecrypt(ciphertext []byte, key []byte) []byte {
	plaintext := []byte{}
	var ki int
	for i := range ciphertext {
		ki = i % len(key)
		plaintext = append(plaintext, ciphertext[i]^key[ki])
	}
	return plaintext
}
