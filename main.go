package main

import (
	"fmt"
	"log"
	"net/http"
)

func GetMessage() (string, error) {
	name := "SLIM"
	return fmt.Sprintf("Welcome to %s's world", name), nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	msg, _ := GetMessage()
	fmt.Fprint(w, msg)
}

func main() {
	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
