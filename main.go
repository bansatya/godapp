package main

import (
	"github.com/bansatya/godapp/app"
)

func main() {

	app := &app.App{}
	app.Initialize()
	app.Run(":10003")
}
