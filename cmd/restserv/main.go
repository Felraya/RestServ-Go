package main

import (
	"fmt"
	"internal/web/rest"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Rest API server with Gorilla Mux Routers")

	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)

	// define routes and handlers
	rest.StudentHandlers(router)
	rest.LanguageHandlers(router)

	// start server but with "router" as second argument
	log.Fatal(http.ListenAndServe(":9999", router)) // log.Fatal = if error : log(error) + exit(1)
}
