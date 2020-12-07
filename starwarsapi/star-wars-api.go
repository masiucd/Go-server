package starwarsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Planet struct
type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
	Gravity    string `json:"gravity"`
}

// Person struct
type Person struct {
	Name         string `json:"name"`
	HomeWorldURL string `json:"homeworld"`
	Gender       string `json:"gender"`
	HomePlanet   Planet
}

// AllPeople struct
type AllPeople struct {
	People []Person `json:"results"`
}

// BaseURL constant
const BaseURL = "https://www.swapi.tech/api/"

func (p *Person) getHomeWorld() {
	res, err := http.Get(p.HomeWorldURL)
	if err != nil {
		log.Print("Failed to fetch home world", err)
	}

	var bytes []byte

	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("error reading planet request", err)
	}

	json.Unmarshal(bytes, &p.HomePlanet)

}

func getPeople(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(BaseURL + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request people")
	}

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("failed to parse request body")
	}

	var people AllPeople

	// similar to json.parse in javascript
	// give me the memory address where we want to store this information
	if err := json.Unmarshal(bytes, &people); err != nil {
		log.Print("Error parsing json", err)
	}

	fmt.Println(people)

	for _, p := range people.People {
		// p.getHomeWorld()
		fmt.Println(p)
	}
}

func main() {
	http.HandleFunc("/", getPeople)
	fmt.Println("server is on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
