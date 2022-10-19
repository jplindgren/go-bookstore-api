package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jplindgren/bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBooksControllerRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server at port: 9010...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
