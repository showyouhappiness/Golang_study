package readWrite

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"testing"
)

func TestHashSha(t *testing.T) {
	hasher := sha1.New()
	io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
	//
	hasher.Reset()
	data := []byte("We shall overcome!")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
	fmt.Printf("Result: %d\n", checksum)
}

func TestHashMD5(t *testing.T) {
	hasher := md5.New()
	b := []byte{}
	io.WriteString(hasher, "test")
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
}
