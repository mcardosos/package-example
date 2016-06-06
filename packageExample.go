package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mcardosos/package-example/goose"
)

var gansho = goose.Goose{
	Eyes:         "cafes",
	Hair:         "cafe",
	Age:          24,
	Intelligence: 10,
	Height:       1.70,
	Weight:       68.5,
	Name:         "Gui"}

func handler(w http.ResponseWriter, r *http.Request) {
	desc := goose.Goose.Describe(gansho)
	fmt.Fprint(w, desc)
}
func main() {
	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
