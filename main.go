package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func welcomeHandler(response http.ResponseWriter, request *http.Request ){
	fmt.Fprintf(response, "Welcome world")
}


var serverPort string = "3030"

func main() {
	fmt.Println("Go String Vowel-Consonant Counter")

	router := mux.NewRouter()
	router.HandleFunc("/", welcomeHandler)
	fmt.Printf("Your server is up on localhost:"+serverPort)

	log.Fatal(http.ListenAndServe(":"+serverPort, router))

}

