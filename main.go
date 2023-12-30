package main

import "github.com/edwincarlflores/mind/server"

func main() {
	err := server.SetupAndRunServer()
	if err != nil {
		panic(err)
	}
}
