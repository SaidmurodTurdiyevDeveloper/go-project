package main

import (
	"fmt"
	"log"
	"net/http"
	"packageinitilization/math"

	"github.com/gorilla/mux"
)

func main() {
	xs := []float64{1, 2, 3, 4}

	avg := math.Average(xs)
	fmt.Println("Average slice of values:", avg)
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHoem).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hello mod in golang")
}

func serveHoem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welocme to golang series on my example</h1>"))
}
