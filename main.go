package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Artist *Artist `json: "artist"`
}

type Artist struct {
	Name string `json:"name"`
}

var albums []Album

func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(albums)
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range albums {
		if item.ID == params["id"] {
			albums = append(albums[:index], albums[index+1:]...)

			break
		}
	}

	json.NewEncoder(w).Encode(albums)
}

func getOneAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range albums {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)

			return
		}
	}
}

func createNewAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var album Album
	_ = json.NewDecoder(r.Body).Decode(&album)

	album.ID = strconv.Itoa(rand.Intn(100000000))
	albums = append(albums, album)
	json.NewEncoder(w).Encode(albums)
}

func updateAlbum(w http.ResponseWriter, r *http.Request) {

	// set json content type
	w.Header().Set("Content-Type", "application/json")

	// params
	params := mux.Vars(r)

	// loop over the albums, range
	for index, item := range albums {
		if item.ID == params["id"] {
			albums = append(albums[:index], albums[index+1:]...)

			var album Album

			// delete the album by id
			_ = json.NewDecoder(r.Body).Decode(&album)
			album.ID = params["id"]

			// set the new body of the album
			albums = append(albums, album)
			json.NewEncoder(w).Encode(album)

			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	albums = append(albums, Album{
		ID:     "1",
		Title:  "Afrociberdelia",
		Year:   1994,
		Artist: &Artist{"Chico Science"},
	})

	albums = append(albums, Album{
		ID:     "2",
		Title:  "Fome de Tudo",
		Year:   2004,
		Artist: &Artist{"Nação Zumbi"},
	})

	albums = append(albums, Album{
		ID:     "3",
		Title:  "Lemonade",
		Year:   2016,
		Artist: &Artist{"Beyoncé"},
	})

	albums = append(albums, Album{
		ID:     "4",
		Title:  "This Hell",
		Year:   2022,
		Artist: &Artist{"Rina Sawayama"},
	})

	albums = append(albums, Album{
		ID:     "5",
		Title:  "Happier Than Ever",
		Year:   2021,
		Artist: &Artist{"Billie Eilish"},
	})

	r.HandleFunc("/albums", getAlbums).Methods("GET")
	r.HandleFunc("/albums/{id}", getOneAlbum).Methods("GET")
	r.HandleFunc("/albums", createNewAlbum).Methods("POST")
	r.HandleFunc("/albums/{id}", updateAlbum).Methods("PUT")
	r.HandleFunc("/albums/{id}", deleteAlbum).Methods("DELETE")

	fmt.Printf("Server is running at port 8080\n")

	log.Fatal(http.ListenAndServe(":8080", r))
}
