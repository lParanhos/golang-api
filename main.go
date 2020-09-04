package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func welcomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo !")
}

func bookList(w http.ResponseWriter, r *http.Request) {
	encode := json.NewEncoder(w)
	encode.Encode(Books)
}

func handlersConfig() {
	http.HandleFunc("/", welcomeRoute)
	http.HandleFunc("/books", bookList)
}

type Book struct {
	ID     int
	Titulo string
	Autor  string
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
