package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s\n\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
