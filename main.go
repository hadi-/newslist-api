package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	ApiKey = " Replace with you newsapi.com KEY"
)

// Display all data by top headline
func GetList(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://newsapi.org/v2/top-headlines?country=id&apiKey=" + ApiKey)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		w.Write(contents)
	}
}

// Search news based on keyword
func GetSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyword := vars["keyword"]
	if keyword == "" {
		keyword = "golang"
	}
	response, err := http.Get("https://newsapi.org/v2/everything?q=" + keyword + "&apiKey=" + ApiKey)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		w.Write(contents)
	}
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetList).Methods("GET")
	router.HandleFunc("/search", GetSearch).Methods("GET")
	router.HandleFunc("/search/{keyword}", GetSearch).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
