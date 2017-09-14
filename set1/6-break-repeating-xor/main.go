package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	contents, _ := ioutil.ReadFile("6.txt")
	spew.Dump(contents)
	ciphertext := make([]byte, base64.StdEncoding.DecodedLen(len(contents)))
	base64.StdEncoding.Decode(ciphertext, contents)
	ks := findKeySize(ciphertext)
	ks = 29

	keyBlocks := transpose(ciphertext, ks)
	var key []byte
	for i := range keyBlocks {
		key = append(key, findKey(keyBlocks[i]))
	}
	fmt.Printf("\nKey size: %d\nKey: \"%s\"\n\n", ks, key)
	spew.Dump(ciphertext)
	plaintext := rotKeyDecrypt(ciphertext, key)
	fmt.Println(string(plaintext))
}
