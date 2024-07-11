package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	fmt.Printf(string(body))
	if err != nil {
		fmt.Println("Error reading the request")

		return
	}
	fmt.Println("Received POST request!")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", postHandler).Methods("POST")

	fmt.Println("Server listenin on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err.Error())
	}

}
