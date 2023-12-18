package main

import (
	"fmt"
	"go_web_apps/pkg/config"
	"go_web_apps/pkg/handlers"
	"go_web_apps/pkg/render"
	"log"
	"net/http"
)

const portnumber = ":8080"



func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server started on port %s", portnumber)
	_ = http.ListenAndServe(portnumber, nil)

}
