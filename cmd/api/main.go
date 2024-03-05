package main

import (
	"log"
	"real_nimi_project/pkg/api"
)

func main() {
	app := api.Default()
	err := app.Start()
	if err != nil {
		log.Print(err)
		panic(err)
	}

}
