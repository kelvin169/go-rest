package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//House Struct(Model)
type House struct {
	ID       string `json:"id"`
	Rooms    string `json:"rooms"`
	Location string `json:"location"`
	Owner    *Owner `json:"owner"`
}

// Author Struct
type Owner struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// init houses var as slice house struct

var houses []House

// get all houses
func getHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get houses

	//Loops through houses and find with id
	for _, item := range houses {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&House{})

}

//get house
func getHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)

}

//create house
func createHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var house House

	_ = json.NewDecoder(r.Body).Decode(&house)
	house.ID = strconv.Itoa(rand.Intn(100000))
	houses = append(houses, house)
	json.NewEncoder(w).Encode(house)

}

//update house
func updateHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range houses {
		if item.ID == params["id"] {
			houses = append(houses[:index], houses[index+1:]...)
			var house House

			_ = json.NewDecoder(r.Body).Decode(&house)
			house.ID = params["id"]
			houses = append(houses, house)
			json.NewEncoder(w).Encode(house)
			return
		}

	}
	json.NewEncoder(w).Encode(houses)
}

func deleteHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range houses {
		if item.ID == params["id"] {
			houses = append(houses[:index], houses[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(houses)
}

func main() {
	// init the mux router
	router := mux.NewRouter()

	//Mock Data @todo implement data
	houses = append(houses, House{ID: "1", Location: "juja", Rooms: "12", Owner: &Owner{Firstname: "kelvin", Lastname: "onkundi"}})
	houses = append(houses, House{ID: "2", Location: "Juja", Rooms: "12", Owner: &Owner{Firstname: "Kelvin", Lastname: "Onkundi"}})

	//create router houndlers/ Endpoings
	router.HandleFunc("/api/houses", getHouses).Methods("GET")
	router.HandleFunc("/api/houses/{id}", getHouse).Methods("GET")
	router.HandleFunc("/api/houses", createHouse).Methods("POST")
	router.HandleFunc("/api/houses/{id}", updateHouse).Methods("PUT")
	router.HandleFunc("/api/houses/{id}", deleteHouse).Methods("DELETE")

	handler := cors.Default().Handler(router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200/"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler = c.Handler(handler)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
