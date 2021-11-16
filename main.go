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

	http.HandleFunc("/", helloHandler)
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
