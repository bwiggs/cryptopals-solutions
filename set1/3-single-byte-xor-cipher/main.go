package main

import (
	"fmt"
	"strings"
)

import "encoding/hex"

func main() {
	in, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	common := strings.Split(strings.ToLower("ETAOIN SHRDLU"), "")

	var best string
	var bestScore int

	var b byte
	for i := 0; i < 255; i++ {
		b = byte(i)
		buf := decrypt(in, b)
		var score int
		for s := range common {
			score += strings.Count(string(buf), common[s])
		}
		if score > bestScore {
			best = fmt.Sprintf("%d\t%x (%s)\t%s\n", score, b, string(b), buf)
			bestScore = score
		}
	}
	fmt.Println(best)
}

func decrypt(s []byte, b byte) (buf []byte) {
	for j := range s {
		buf = append(buf, s[j]^b)
	}
	return
}
