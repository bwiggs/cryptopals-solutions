package main

import "fmt"
import "os"
import "encoding/hex"

func main() {
	s1 := os.Args[1]
	s2 := os.Args[2]
	x, _ := hex.DecodeString(s1)
	y, _ := hex.DecodeString(s2)

	z := []byte{}
	for i := range x {
		z = append(z, x[i]^y[i])
	}

	fmt.Printf("%s", hex.EncodeToString(z))
}
