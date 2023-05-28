package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/file", UploadFile).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

type A struct {
	Id    string
	Index int
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	fileName := r.FormValue("file_name")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = io.WriteString(w, "File "+fileName+" Upload file succesfully")
	if err != nil {
		log.Fatal(err)
	}
	dir, err := f.ReadDir(-1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	_, err = io.Copy(f, file)
	if err != nil {
		log.Fatal(err)
	}
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome home!")
	if err != nil {
		log.Fatal(err)
	}
}
