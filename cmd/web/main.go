package main

import (
	"fmt"
	"net/http"
	"go_web_apps/pkg/handlers"
)

const portnumber = ":8080"



func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server started on port %s", portnumber)
	_ = http.ListenAndServe(portnumber, nil)

}
