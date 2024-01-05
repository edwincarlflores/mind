package main

import "github.com/edwincarlflores/mind/internal/app"

func main() {
	err := app.SetupAndRun()
	if err != nil {
		panic(err)
	}
}
