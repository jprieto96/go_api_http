package main

import (
	"log"

	"github.com/jprieto96/go_api_http/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
