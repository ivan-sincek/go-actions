package goactions

import (
	"fmt"
	"crypto/md5"
	"testing"
)

func TestCrypto(t *testing.T) {
	data := []byte("Hello World!")
	fmt.Printf("%x\n", md5.Sum(data))
}

func TestBasic(t *testing.T) {
	var password = "123456789"
	if password == "123456789" {
		fmt.Print("Hello World!\n")
	}
}
