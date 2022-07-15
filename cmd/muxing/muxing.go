package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"io/ioutil"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

func NameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hello, %s!", vars["param"])
}

func BadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "I got message:\n%s", string(body))
	}
}

func HeadersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		a := r.Header.Get("a")
		b := r.Header.Get("b")

		a2, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		b2, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}
		ab := strconv.Itoa(a2 + b2)

		w.Header().Set("a+b", ab)
	}
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{param}", NameHandler)
	router.HandleFunc("/bad", BadHandler)
	router.HandleFunc("/data", DataHandler)
	router.HandleFunc("/headers", HeadersHandler)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
