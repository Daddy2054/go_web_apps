package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			fmt.Println(err)

		}
		var a []any = []any{fmt.Sprintf("Number of bytes written %d", n)}
		fmt.Fprintln(os.Stdout, a...)
	})

	_ = http.ListenAndServe(":8080", nil)

}
