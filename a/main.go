package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not find", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not allowed", http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello")
}
func form(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm()  %v", err)
		return
	}
	fmt.Fprintf(w, "POST SUCCESS")
	name := r.FormValue("name")
	password := r.FormValue("password")
	fmt.Fprintf(w, "%v:%v", name, password)

}
func main() {
	server := http.FileServer((http.Dir("./static")))
	http.Handle("/", server)
	http.HandleFunc("/form", form)
	http.HandleFunc("/hello", hello)
	fmt.Printf("start server")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
