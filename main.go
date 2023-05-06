package main

import (
	"fmt"
	"log"

	"github.com/just-nibble/dockerish/linux"
)

func main() {
	ids, err := linux.GetAllContainerIDs()
	if err != nil {
		log.Fatal(err)
	}
	message, err := linux.StopAllContainers(ids)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
