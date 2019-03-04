package main

import (
	"log"

	"golang.org/x/crypto/sha3"
)

func main() {
	log.Println("Running byhash utility")

	data := "hello world"
	hash := toKeccak([]byte(data))

	log.Printf("%x\n", hash)
}

func toKeccak(data []byte) []byte {
	hash := make([]byte, 64)

	keccak := sha3.NewLegacyKeccak256()
	keccak.Write(data)
	hash = keccak.Sum(nil)

	return hash
}
