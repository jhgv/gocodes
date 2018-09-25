package main

import (
	"log"
	"net/http"

	"github.com/jhgv/gocodes/blog/settings"
)

func main() {
	settings.MapUrls()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
