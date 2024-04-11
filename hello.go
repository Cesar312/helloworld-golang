package main

import (
	"log"
	"net/http"
	"time"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	responseMessage := "Hello World! Current Date and Time: " + currentTime
	w.Write([]byte(responseMessage))
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
