package viewResolver

import (
	"html/template"
	"net/http"

	"github.com/krix38/ScorchedGo/properties"
)

var Templates = template.Must(template.ParseGlob(properties.Configuration.TemplateFilesPath + "/*"))

func Render(w http.ResponseWriter, templateName string, model interface{}) {
	err := Templates.ExecuteTemplate(w, templateName+".html", model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
