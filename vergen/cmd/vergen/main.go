package main

import (
	"log"

	"github.com/Akagi201/utils-go/vergen"
)

func main() {
	err := vergen.Create()
	if err != nil {
		log.Fatalln(err)
	}
}
