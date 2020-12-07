package main

import (
	"fmt"
	"log"
	"net/http"
)

const baseUrl = "https://swapi.co/api/"

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func main() {
	fmt.Println(baseUrl)
	http.HandleFunc("/", home)
	fmt.Println("server is on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
