package main

import (
	"ScrumPokerBackend/app"



)

func main() {
	app := app.App{}
	app.Initialize()
	app.Run(":8000")
}
