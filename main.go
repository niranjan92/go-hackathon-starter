package main

import (
	"log"

	"github.com/niranjan92/go-hackathon-starter/actions/app"
)

func main() {
	app := app.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
