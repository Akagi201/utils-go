package main

import (
	"fmt"
	"time"

	"github.com/Akagi201/utilgo/broadcast"
)

func main() {
	b := broadcast.NewBroadcast()

	go func() {
		for {
			b.Send(1)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println(<-b.Receive())
			fmt.Println("2")
		}
	}()

	for {
		fmt.Println(<-b.Receive())
	}

	select {}
}
