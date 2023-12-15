package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portnumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum:=addValues(2,2)
	_,_=fmt.Fprintf(w, "This is the about page and the sum  is %d", sum)
	// _,_=fmt.Fprintf(w, fmt.Sprintf("This is the about page and the sum of 2 +2 is %d", sum))
}

func addValues(x, y int) (int)  {

	return x + y
}



func divideValues(x,y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	} else {
		return x / y, nil
	}
}

func Divide(w http.ResponseWriter, r *http.Request){
	f,err:=divideValues(100.0,0.0)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")
		return
	}
	fmt.Fprintf(w, "%f divided by %f is %f", 100.0,0.0,f)
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Server started on port %s", portnumber)
	// fmt.Printf(fmt.Sprintf("Server started on port %s", portnumber))
	_ = http.ListenAndServe(portnumber, nil)

}
