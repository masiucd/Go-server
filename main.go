package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/masiuciszek/go-server/routes"
)




func main() {
	http.HandleFunc("/", routes.Home)	
	http.HandleFunc("/users", routes.Users)
	http.HandleFunc("/tasks", routes.Tasks)
	http.HandleFunc("/hello", routes.Hello)
	fmt.Println("server is on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
