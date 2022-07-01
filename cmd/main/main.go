package main

import (
	"log"
	"net/http"

	"github.com/Threx-code/go-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RequestHandler(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9080", router))

}
