package main

import (
		"encoding/json"
    "time"
    "fmt"
    "log"
    "github.com/gorilla/mux"
    "net/http"
)
type Doc struct {
    Time      time.Time   `json:"t_stamp"`
    Pos       int         `json:"pos"`
    Content   string      `json:"content"`

}
type Monolyth []Doc

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Monolyth{
		Doc{Content: "write presentation"},
		Doc{Content: "host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
	//fmt.Fprintln(w, "TODO index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
