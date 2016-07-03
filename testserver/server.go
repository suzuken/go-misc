package main

import (
	"fmt"
	"log"
	"net/http"
)

// Routeはこのアプリケーションのルーティングを設定しています
// http.ServeMuxを返します
func Route() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Fprintf(w, "Hello, %s", r.FormValue("name"))
	})
	return m
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", Route()))
}
