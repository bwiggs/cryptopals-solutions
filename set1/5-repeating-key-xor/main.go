package main

import "fmt"

import "encoding/hex"

func main() {
	line := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := []byte("ICE")
	ki := 0
	fmt.Println(line)
	o := []byte{}
	for _, ch := range line {
		o = append(o, byte(ch)^key[ki])
		if ki == 2 {
			ki = 0
		} else {
			ki++
		}
	}
	fmt.Println(hex.EncodeToString(o))
}
