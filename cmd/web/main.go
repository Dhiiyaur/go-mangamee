package main

import (
	"github.com/dhiiyaur/go-mangamee/internal/app"
	"github.com/joho/godotenv"
)

func main() {

	initEnvConfiguration()

	app.Start()
}

func initEnvConfiguration() {
	godotenv.Load(".env")
}
