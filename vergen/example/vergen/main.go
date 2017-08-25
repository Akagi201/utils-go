package main

import (
	"log"

	"github.com/Akagi201/utilgo/vergen"
)

func main() {
	err := vergen.Create()
	if err != nil {
		log.Fatalln(err)
	}
}
