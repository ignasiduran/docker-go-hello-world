package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "HEAD" && r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	t := time.Now()
	fmt.Fprintln(w, "Hello World! It's "+t.Format("2006-01-02 15:04:05"))
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
}

func healthy(w http.ResponseWriter, r *http.Request) {
	if r.Method != "HEAD" && r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "Healthy")
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
}

func main() {

	http.HandleFunc("/healthy", healthy)
	http.HandleFunc("/", index)

	log.Println("	* Http Server started")
	log.Println("	* Running on http://0.0.0.0:8080/ (Press CTRL+C to quit)")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
