package webapp

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/incidrthreat/GoAccomplish/models"
)

// App ...
type App struct {
	HTMLDir   string
	StaticDir string
	Database  *models.Dataservice
	Store     sessions.CookieStore
}

func (app *App) RenderHTML(w http.ResponseWriter, r *http.Request, page string, data interface{}) {
	files := []string{
		filepath.Join(app.HTMLDir, "base.html"),
		filepath.Join(app.HTMLDir, page),
	}
	ts, _ := template.ParseFiles(files...)

	//if err != nil {
	//	app.ServerError(w, err) // Use the new app.ServerError() helper.
	//	return
	//}

	ts.ExecuteTemplate(w, "base", nil)
}
