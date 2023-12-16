package main

import (
	"fmt"
	"net/http"
)

const portnumber = ":8081"



func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server started on port %s", portnumber)
	_ = http.ListenAndServe(portnumber, nil)

}
