package main

import (
	"log"

	"github.com/niranjan92/go-hackathon-starter/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
