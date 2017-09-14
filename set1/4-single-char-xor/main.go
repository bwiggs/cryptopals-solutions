package main

import "fmt"
import "strings"
import "io/ioutil"

import "encoding/hex"

func main() {
	content, _ := ioutil.ReadFile("4.txt")
	lines := strings.Split(string(content), "\n")
	vowels := strings.Split("etaoin shrdlu", "")
	var bestScore int
	var plaintext string
	for _, s := range lines {
		in, _ := hex.DecodeString(s)
		for i := 0; i < 255; i++ {
			b := byte(i)
			buf := decrypt(in, b)
			var score int
			for s := range vowels {
				score += strings.Count(string(buf), vowels[s])
			}
			if score > bestScore {
				plaintext = fmt.Sprintf("%d - %x - %s\n", score, b, buf)
				bestScore = score
			}
		}
	}
	fmt.Println(plaintext)
}

func decrypt(s []byte, b byte) (buf []byte) {
	for j := range s {
		buf = append(buf, s[j]^b)
	}
	return
}
