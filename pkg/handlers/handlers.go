package handlers

import (
	"net/http"

	"github.com/JOFosuware/bookings/pkg/config"
	"github.com/JOFosuware/bookings/pkg/models"
	"github.com/JOFosuware/bookings/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the respository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandler sets the respository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.gtpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	//send the data to template
	render.RenderTemplate(w, "about.page.gtpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
