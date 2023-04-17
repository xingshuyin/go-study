package main

import (
	"c/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.Init_book_route(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))

}

// xapp-1-A04KHABCR5K-4651769523079-8e62b4d1ad2d923f6c1e43d4601b8c28a444455e265a5b37ba655eff028bd34e
