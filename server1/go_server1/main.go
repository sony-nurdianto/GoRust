package main

import (
	"go_server1/lissajous"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous.Lissajous(w)
	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
