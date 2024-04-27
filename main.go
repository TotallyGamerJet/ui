package main

import (
	"log"
	"net/http"
	"ui/components"
)

func main() {
	http.DefaultServeMux.Handle("/output.css", http.FileServer(http.Dir("./src")))
	http.DefaultServeMux.HandleFunc("/", components.Index)

	log.Fatalln(http.ListenAndServe(":8080", http.DefaultServeMux))
}
