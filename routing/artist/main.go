package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Artist struct {
	id   string
	name string
}

func New(id, name string) *Artist {
	return &Artist{id: id, name: name}
}

func main() {
	r := mux.NewRouter()
	artists := make([]Artist, 0)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome to artist management!")
	})

	artistRouter := r.PathPrefix("/artist").Subrouter()

	/** @Get artist **/
	artistRouter.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

		var found bool
		var foundArtist Artist

		for _, v := range artists {
			if v.name == name {
				found = true
				foundArtist = v
				break
			}
		}

		if found {
			fmt.Fprintf(w, "success, artist found, %v", foundArtist.name)
		} else {
			fmt.Fprintf(w, "sorry, artist not found")
		}

	}).Methods("GET")

	/** @Create new artist **/
	artistRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		name := r.FormValue("name")

		if id == "" || name == "" {
			fmt.Fprintf(w, "id and name fields are required")
			return
		}

		newArtist := New(id, name)
		artists = append(artists, *newArtist)

		fmt.Fprintf(w, "success add artist")
	}).Methods("POST")

	/** @Delete artist **/
	artistRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var updateArtist []Artist
		var foundArtistID bool

		for _, v := range artists {
			if v.id != id {
				updateArtist = append(updateArtist, v)
			} else {
				foundArtistID = true
			}
		}

		if foundArtistID {
			artists = updateArtist
			fmt.Fprintf(w, "artist success deleted")
		} else {
			fmt.Fprintf(w, "sorry, failed deleted artist, artist with id :%v not found.", id)
		}

	}).Methods("DELETE")

	http.ListenAndServe("localhost:3000", r)
}
