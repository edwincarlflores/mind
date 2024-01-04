package main

import web "github.com/edwincarlflores/mind/internal/adapters/web"

func main() {
	err := web.SetupAndRunWebServer()
	if err != nil {
		panic(err)
	}
}
