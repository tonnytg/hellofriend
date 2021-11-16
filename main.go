package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func StartAPI() {

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello Friend!\n")
	}

	helloHandlerCloud := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello Cloud Engineer!\n")
	}

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/cloud", helloHandlerCloud)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	log.Println("StartAPI")
	go StartAPI()

	for {
		log.Println("Hello Friend")
		time.Sleep(time.Second * 5)
	}
}
