package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
func TaskManager(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Task manager")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/tasks", TaskManager)

	fmt.Println("server is on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
