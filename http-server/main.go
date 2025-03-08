package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("hi")
		param = strings.Trim(param, "\"")

		fmt.Println(param)

		fmt.Fprintln(w, "Hello world")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	http.ListenAndServe("localhost:3000", nil)
}
