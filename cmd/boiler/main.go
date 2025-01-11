package main

import (
	"github.com/joho/godotenv"
	app "github.com/placeholder/boiler/app"
)

func main() {
	godotenv.Load()
	app.Run()
}
