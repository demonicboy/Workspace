package handlers

import (
	"net/http"
	"workspace/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "This is the Home page")
	render.RenderTemplate(w, "home.page.tmpl")
}
func About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "This is the About page")
	render.RenderTemplate(w, "about.page.tmpl")
}
