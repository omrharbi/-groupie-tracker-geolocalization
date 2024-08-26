package main

import (
	"fmt"
	"log"
	"net/http"

	Groupie_tracker "groupie_tracker/Funcs"
)

func main() {
	port := ":8093"
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))
	http.HandleFunc("/", Groupie_tracker.GetDataFromJson)
	http.HandleFunc("/Artist/{id}", Groupie_tracker.HandlerShowRelation)
	http.HandleFunc("/styles/", Groupie_tracker.HandleStyle)
	http.HandleFunc("/maps/", Groupie_tracker.Handler)

	fmt.Printf("http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
