package routes

import (
	"fmt"
	"net/http"
	"time"
)

// Home route
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

// Users route
func Users(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Users")
}

// Tasks route
func Tasks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Tasks")
}

// Hello route
func Hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "Hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)

	}

}
