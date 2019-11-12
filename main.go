package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func welcomeHandler(response http.ResponseWriter, request *http.Request){
	fmt.Fprintf(response, "Welcome to Vowel-Consonant Counter!\n\nFor usage you can add paths:\n- /count/{YourStringWithoutCurlyBracket}\n")
}

func countHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	myString := vars["userString"] //get parameter then assign to mystring
	fmt.Println(myString)
	
	vowel := 0
	conso := 0
	// var savedWord rune
	var filteredString []rune
	runes := []rune(myString)

	// remove duplicate word
	for i := 0; i < len(runes); i++ {
		// Scan slice for a previous element of the same value.
		exists := false
		for v := 0; v < i; v++ {
			if runes[v] == runes[i] {
				exists = true
				break
			}
		}
		// If no previous element exists, append this one.
		if !exists {
			filteredString = append(filteredString, runes[i])
		}
	}

	for _, value := range filteredString {
		switch value {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowel++
		default :
			conso++
		}
		
	}
	response.WriteHeader(http.StatusOK)
	fmt.Printf("%d Vowels, %d Consonant \n", vowel,conso)
	response.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(response, "%d Vowels, %d Consonant", vowel,conso)
}



var serverPort string = "3030"

func main() {
	fmt.Println("Go String Vowel-Consonant Counter")

	router := mux.NewRouter()
	router.HandleFunc("/", welcomeHandler)
	router.HandleFunc("/count/{userString}", countHandler)
	fmt.Printf("Your server is up on localhost:"+serverPort)

	log.Fatal(http.ListenAndServe(":"+serverPort, router))

}


func readLine(message string) string{
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}