package main

import (
	"fmt"
	"net/http"
)

func handlersConfig() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bem vindo !")
	})

}

func serverConfig() {
	handlersConfig()

	fmt.Println("Server is listening in port 1337....")
	http.ListenAndServe(":1337", nil) // DefaultServerMux
}

func main() {
	serverConfig()
}
