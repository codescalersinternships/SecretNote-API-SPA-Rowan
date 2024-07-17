package main

import (
	"log"

	pkg "github.com/codescalersinternships/SecretNote-API-SPA-Rowan/pkg"
)

func main() {
	app := pkg.NewApp()
	if err := app.Run(); err != nil {
		log.Fatal("server isn't starting")
	}
}
