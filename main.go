package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portnumber = ":8081"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.go.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.go.html")

}
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server started on port %s", portnumber)
	_ = http.ListenAndServe(portnumber, nil)

}
