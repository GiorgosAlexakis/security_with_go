package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	s := "hello\n"
	h := md5.New()
	h.Write([]byte(s))
	hashValue := h.Sum(nil)
	fmt.Printf("Hash value in binary : %b\n", hashValue)
	fmt.Printf("Hash value in decimal : %d\n", hashValue)
	fmt.Printf("Hash value in hexadecimal : %x\n", hashValue)
	hashSize := h.Size()
	fmt.Printf("Size: %d bytes\n", hashSize)
	fmt.Println("Usual form is 4 hexadecimal numbers (4 bytes each):")
	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n])<<24 +
			uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 +
			uint32(hashValue[n+3])
		fmt.Printf("%x ", val)
	}
	fmt.Println()
}
