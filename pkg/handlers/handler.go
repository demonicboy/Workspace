package handlers

import (
	"net/http"
	"workspace/pkg/config"
	"workspace/pkg/models"
	"workspace/pkg/render"
)

// the repository used by the handlers
var Repo *Repository

// the repository type
type Repository struct {
	App *config.AppConfig
}

// create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "This is the Home page")
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	//fmt.Fprintf(w, "This is the About page")
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
