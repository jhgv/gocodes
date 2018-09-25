package urls

import (
	"net/http"

	"github.com/jhgv/gocodes/blog/handlers"
)

// MapUrls : Function that mapps the handlers with URLs
func MapUrls() {
	http.HandleFunc("/", handlers.RedirectHome)
	http.HandleFunc("/home/", handlers.MakeHandler(handlers.Home))
	http.HandleFunc("/save/", handlers.MakeHandler(handlers.Save))
	http.HandleFunc("/edit/", handlers.MakeHandler(handlers.Edit))
	http.HandleFunc("/view/", handlers.MakeHandler(handlers.View))
}
