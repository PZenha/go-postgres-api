package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type IoTValue struct {
	Temp     float64 `json:"Temp"`
	Humidity float64 `json:"Humidity"`
	AC       bool    `json:"AC"`
}

type IoTValues []IoTValue

func setRoomValues(w http.ResponseWriter, r *http.Request) {
	//decoder := json.NewDecoder(r.Body)
	temp := r.Form.Get("Temp")
	fmt.Printf("temp: %s", temp)
	fmt.Fprintf(w, "Temp: %s", temp)
}

func getRoomValues(w http.ResponseWriter, r *http.Request) {
	values := IoTValues{
		IoTValue{Temp: 23.2, Humidity: 60.5, AC: true},
	}
	json.NewEncoder(w).Encode(values)
}

func handleRootPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "am alive maite!")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/room", getRoomValues)
	router.HandleFunc("/send", setRoomValues).Methods("POST")
	router.HandleFunc("/", handleRootPath)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
