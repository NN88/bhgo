package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hw", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})
	http.Handle("/", http.FileServer(http.Dir("./static")))

	port := ":5000"
	fmt.Printf("Server is running on port %v", port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))

	// Then you can
	// curl http://localhost:5000
	// curl http://localhost:5000/hi

}
