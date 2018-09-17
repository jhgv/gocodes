package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jhgv/gocodes/blog/models"
	"github.com/jhgv/gocodes/blog/settings/templates"
)

func loadPage(title string) (*models.Page, error) {
	filename := title + ".txt"
	log.Print(filename)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &models.Page{Title: title, Body: body}, nil
}

// RedirectHome : handler to redirect web root to home screen
func RedirectHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home/", http.StatusFound)
}

// Home : handler for home screen
func Home(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage("")
	if err != nil {
		p = &models.Page{Title: "home"}
	}
	templates.RenderTemplate(w, "index", p)
}

// Edit : handler to edit a blog posts
func Edit(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &models.Page{Title: title}
	}
	templates.RenderTemplate(w, "edit", p)
}

// View : handler to view a blog posts
func View(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	templates.RenderTemplate(w, "view", p)
}

// Save : handler to save a blog posts
func Save(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &models.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
