package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}
type TodoPageData struct {
	Metadata  string
	PageTitle string
	Todos     []Todo
}
type AboutPageData struct {
	Metadata string
	About    string
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))
	aboutTmpl := template.Must(template.ParseFiles("about.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			Metadata:  "This is server side rendering",
			PageTitle: "My Todo List",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
				{Title: "Task 4", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := AboutPageData{
			Metadata: "About me",
			About:    "my name is fais, just a guy who trying to fit in to the world :). Big fan of lizzy mcalpine!",
		}
		aboutTmpl.Execute(w, data)
	})

	http.ListenAndServe("localhost:3000", nil)
}
