package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/yjunya/api-example/cmd/app"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	app.Run()
}
