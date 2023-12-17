package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenTemplate(){}
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl,"./templates/base.layout.go.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}