package main

import (
	"github.com/eddievagabond/internal/application"
)

func main() {
	app := application.Initialize()
	app.Run()
}
