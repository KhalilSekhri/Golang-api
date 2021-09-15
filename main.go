package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Hello pas world enft")
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}
		for key, value := range req.PostForm {
			fmt.Println(key, "=>", value)
		}
		fmt.Fprintf(w, "Information received: %v\n", req.PostForm)
	}
}
func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":9000", nil)
}
