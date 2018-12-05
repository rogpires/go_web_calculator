package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Ola WEB! </h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
