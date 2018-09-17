package templates

import (
	"html/template"
	"net/http"
	"path"

	"github.com/jhgv/gocodes/blog/models"
)

var templatePath = path.Join("templates", "*")

var templates = template.Must(template.ParseGlob(templatePath))

// RenderTemplate : Function that reads template from template cache
func RenderTemplate(w http.ResponseWriter, templateName string, p *models.Page) {
	err := templates.ExecuteTemplate(w, templateName+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
