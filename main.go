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
	t := 0
	c := 0


	for _, value := range myString {
		switch value {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			t++
		default :
			c++
		}
		
	}
	response.WriteHeader(http.StatusOK)
	fmt.Printf("/n %d Vowels, %d Consonant", t,c)
	response.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(response, "%d Vowels, %d Consonant", t,c)
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