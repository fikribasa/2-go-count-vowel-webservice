
// endpoints_test.go
package main

import (
	"fmt"
    "net/http"
    // "net/http/httptest"
	"testing"
	
	// "github.com/gorilla/mux"
)

// func TestWelcomeHandler(t *testing.T) {
//     // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
//     // pass 'nil' as the third parameter.
//     req, err := http.NewRequest("GET", "/", nil)
//     if err != nil {
//         t.Fatal(err)
//     }

//     // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
//     rr := httptest.NewRecorder()
//     handler := http.HandlerFunc(welcomeHandler)

//     // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
//     // directly and pass in our Request and ResponseRecorder.
//     handler.ServeHTTP(rr, req)

//     // Check the status code is what we expect.
//     if status := rr.Code; status != http.StatusOK {
//         t.Errorf("handler returned wrong status code: got %v want %v /n",
//             status, http.StatusOK)
//     }

// }

func TestCountHandler(t *testing.T) {
    tt := []struct{
        userString string
		vowel int
		conso int
    }{
        {"goroutines", 4,5},
        {"heap", 2,2},
        {"counters", 3,5},
        {"queries", 3,3},
        {"adhadaeqm3k", 2,6},
    }

    for _, tc := range tt {
        path := fmt.Sprintf("/count/%s", tc.userString)
        _, err := http.NewRequest("GET", path, nil)
        if err != nil {
            t.Fatal(err)
        }

    //     rr := httptest.NewRecorder()
	
	// // Need to create a router that we can pass the request through so that the vars will be added to the context
	// router := mux.NewRouter()
    //     router.HandleFunc("/count/{userString}", countHandler)
    //     router.ServeHTTP(rr, req)

	// 	fmt.Println(http.response)
	// 	// if s != val.output1 || k != val.output2 {
	// 	// 	t.Errorf("countHandler(%s) => %d, %d Should be %d, %d /n", val.input, s, k, val.output1, val.output2)
	// 	// }
    //     // In this case, our MetricsHandler returns a non-200 response
    //     // for a route variable it doesn't know about.
    //     if rr.Code == http.StatusOK  {
    //         t.Errorf("handler should have failed on userString %s: got %v want %v",
    //             tc.userString, rr.Code, http.StatusOK)
    //     }
    }
}