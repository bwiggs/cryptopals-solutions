package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
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

	// GoLang doesn't ship with an AES ECB cipher since it's insecure.
	// see: https://github.com/golang/go/issues/5597

	numEntries := len(ciphertext) / aes.BlockSize
	str := ""
	for i := 0; i < numEntries; i++ {
		dst := make([]byte, aes.BlockSize)
		block.Decrypt(dst, ciphertext[aes.BlockSize*i:])
		str += string(dst)
	}
	fmt.Println(str)
}
