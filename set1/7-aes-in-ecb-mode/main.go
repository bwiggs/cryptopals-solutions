package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	key := []byte("YELLOW SUBMARINE")
	contents, _ := ioutil.ReadFile("7.txt")
	ciphertext := make([]byte, base64.StdEncoding.DecodedLen(len(contents)))
	base64.StdEncoding.Decode(ciphertext, contents)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	numEntries := len(ciphertext) / aes.BlockSize
	spew.Dump(numEntries)

	str := ""
	for i := 0; i < numEntries; i++ {
		dst := make([]byte, aes.BlockSize)
		block.Decrypt(dst, ciphertext[aes.BlockSize*i:])
		str += string(dst)
	}
	fmt.Println(str)
}
