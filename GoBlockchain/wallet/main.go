package main

import (
	"GoBlockchain/wallet"
	"fmt"
	"log"
)

// next step is create a transaction and adding them into pool.

func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr)
	fmt.Println(w.PublicKeyStr)

}
