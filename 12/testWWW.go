package main

import (
	"fmt"
	"net/http"
	"os"
)

func CheckStatusOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "fine!")
}

func StatusNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	PORT := "8001"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Using default port number: ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}

	http.HandleFunc("/CheckStatusOK", CheckStatusOK)
	http.HandleFunc("/StatusNotFound", StatusNotFound)
	http.HandleFunc("/", myHandler)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}