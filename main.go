package main

import (
	"fmt"
	"net/http"
)

const portnumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum:=addValues(2,2)
	_,_=fmt.Fprintf(w, fmt.Sprintf("This is the about page and the sum of 2 +2 is %d", sum))
}

func addValues(x, y int) (int)  {

	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf(fmt.Sprintf("Server started on port %s", portnumber))
	_ = http.ListenAndServe(portnumber, nil)

}
