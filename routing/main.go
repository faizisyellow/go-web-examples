package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome!")
	})

	r.HandleFunc("/artist/{name}/songs/{song}", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("trending")
		vars := mux.Vars(r)
		name := vars["name"]
		songs := vars["song"]

		fmt.Println(vars)
		fmt.Printf("param: %v, type: %T\n", param, param)

		fmt.Fprintf(w, "it's %v!\n her song is %v", name, songs)
	})

	moviesRouter := r.PathPrefix("/movies").Subrouter()
	moviesRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "get movies!")
	})
	moviesRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "get movies!")
	})
	moviesRouter.HandleFunc("/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]

		fmt.Fprintf(w, "get movies! with title: %v", title)
	})

	http.ListenAndServe("localhost:3000", r)
}
