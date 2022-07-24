package goactions

import (
	"fmt"
	"crypto/md5"
)

func Crypto() {
	data := []byte("Hello World!")
	fmt.Printf("%x\n", md5.Sum(data))
}

func Basic() {
	var password = "123456789"
	if password == "123456789" {
		fmt.Println("Hello World!")
	}
}

func main() {
	var unused int = 3301
	Basic()
	Crypto()
}
