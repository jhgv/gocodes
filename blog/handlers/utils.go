package handlers

import (
	"log"
	"net/http"
	"regexp"
)

var validPath *regexp.Regexp

func mountValidPath() {
	var pathRegex string
	pathRegex += "^("
	pathRegex += "/home/|"
	pathRegex += "/save/([a-zA-Z0-9]+)|"
	pathRegex += "/view/([a-zA-Z0-9]+)|"
	pathRegex += "/edit/([a-zA-Z0-9]+)"
	pathRegex += ")$"
	log.Printf("regex: %s", pathRegex)
	validPath = regexp.MustCompile(pathRegex)
}

// MakeHandler :
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path: %s", r.URL.Path)
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		log.Println(m[2])
		fn(w, r, m[2])
	}
}
