package main

import (
	"log"
	"net/http"

	"github.com/jhgv/gocodes/blog/settings/urls"
)

func main() {
	urls.MapUrls()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
