package templates

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// RenderTemplate renders the template tmpl using data.
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	cwd := os.Getenv("ADMINPORTAL_TEMPLATE_BASEDIR")
	if cwd == "" {
		cwd, _ = os.Getwd()
	}
	t, err := template.ParseFiles(filepath.Join(cwd, "./routes/"+tmpl+"/"+tmpl+".html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
