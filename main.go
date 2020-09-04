package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func welcomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo !")
}

func bookList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encode := json.NewEncoder(w)
	encode.Encode(Books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		//error handler
	}

	var newBook Book
	json.Unmarshal(body, &newBook)
	newBook.ID = len(Books) + 1
	Books = append(Books, newBook)

	encode := json.NewEncoder(w)
	encode.Encode(Books)
}

func booksRoutes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		bookList(w, r)
	} else if r.Method == "POST" {
		createBook(w, r)
	}
}

func handlersConfig() {
	http.HandleFunc("/", welcomeRoute)
	http.HandleFunc("/books", booksRoutes)
}

type Book struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

var Books []Book = []Book{
	Book{
		ID:     1,
		Titulo: "Eric Clapton",
		Autor:  "Eric",
	},
	Book{
		ID:     2,
		Titulo: "Cazuza",
		Autor:  "Viriato Correia",
	},
	Book{
		ID:     3,
		Titulo: "Dom Casmurro",
		Autor:  "Machado de Assis",
	},
}

func serverConfig() {
	handlersConfig()

	fmt.Println("Server is listening in port 1337....")
	log.Fatal(http.ListenAndServe(":1337", nil)) // DefaultServerMux
}

func main() {
	serverConfig()
}
