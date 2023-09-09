package main

import (
	"github.com/edwincarlflores/mind/app"
)

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
