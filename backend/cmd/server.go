package main

import (
	"log"

	pkg "github.com/codescalersinternships/SecretNote-API-SPA-Rowan/pkg"
)

func main() {
	app, err := pkg.NewApp()
	if err != nil {
		log.Fatal("db migration error probably")
	}
	if err := app.Run(); err != nil {
		log.Fatal("server isn't starting")
	}
}
