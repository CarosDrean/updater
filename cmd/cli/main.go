package main

import (
	"log"

	"github.com/CarosDrean/updater/cmd/cli/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
