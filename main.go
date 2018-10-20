package main

import (
	"net/http"
	"os"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}
	AddHandlers()

	http.ListenAndServe(":"+PORT, nil)
}
