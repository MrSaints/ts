package main

import (
	"flag"
	"fmt"
	"github.com/keybase/go-triplesec"
	"io/ioutil"
	"log"
	"os"
)

func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	pw := flag.String("p", "password", "a password")
	d := flag.Bool("decrypt", false, "decryption mode")
	flag.Parse()

	c, err := triplesec.NewCipher([]byte(*pw), nil)
	FailOnError(err)

	var cipher []byte

	if *d {
		log.Println("[TripleSec] Decrypting...")
		cipher, err = c.Decrypt(b)
	} else {
		log.Println("[TripleSec] Encrypting...")
		cipher, err = c.Encrypt(b)
	}
	FailOnError(err)

	fmt.Printf("%s", cipher)
}
