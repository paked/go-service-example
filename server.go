package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("http request. #2")
		fmt.Fprintln(w, "Hello, World!")
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
