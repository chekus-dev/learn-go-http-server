package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/", rootHandler)
	fmt.Println("starting the server on port :3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) { // Haandles request to /home
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return

	}
	fmt.Fprintln(w, "Welcome")

}

func rootHandler(w http.ResponseWriter, r *http.Request) { // Handles request to the root path (/)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	fmt.Fprintln(w, "This is the root page")
}
