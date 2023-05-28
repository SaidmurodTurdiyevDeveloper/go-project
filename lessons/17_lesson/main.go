package main

import (
	"example/go/mongodb/router"
	"log"
	"net/http"
)

func main() {
	r := router.MyRouter()
	log.Fatal(http.ListenAndServe(":4000", r))

}
