// +build development
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gobuffalo/envy"
)

func init() {
	go func() {
		fmt.Println("inside")
		if envy.Get("GO_ENV", "development") == "development" {
			fmt.Println("start pprof on :6060")
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}
	}()
}
