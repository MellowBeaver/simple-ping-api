package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/arithmetic", pingHandler)

	http.ListenAndServe(":8080", nil)
}

type Solution struct {
	//value string `json:"Buffer,omitempty"``
	// always use back quote ``

	Result int `json:"result"`
}

func pingHandler(w http.ResponseWriter, req *http.Request) {

	//fmt.Fprintf(w, "Hello world")

	//cheking for GET req
	if req.Method == "GET" {

		fmt.Fprintln(w, "Second API")

		x := req.URL.Query().Get("x")
		y := req.URL.Query().Get("y")

		i, _ := strconv.Atoi(x)
		j, _ := strconv.Atoi(y)

		sum := i + j

		//storing into struct
		a := Solution{Result: sum}
		aBytes, err := json.Marshal(a)

		log.Println(err)
		log.Println(string(aBytes))

		fmt.Fprintf(w, string(aBytes))

		//perform addition
		//fmt.Fprintf(w, "Answer = %v", sum)

	} else if req.Method == "DELETE" {
		//check for DELETE req
		// Getting data from DELETE request body
		decoder := json.NewDecoder(req.Body)

		// Structure for the DELETE request json body
		type test_struct3 struct {
			X int
			Y int
		}

		var t test_struct3

		// Decoding JSON into test struct
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}

		diff := t.X - t.Y

		//storing into struct
		a := Solution{Result: diff}
		aBytes, err := json.Marshal(a)

		log.Println(err)
		log.Println(string(aBytes))

		fmt.Fprintf(w, string(aBytes))

		//perform subtraction
		//fmt.Fprintf(w, "Answer = %v", diff)

	} else if req.Method == "PUT" {
		//check for PUT req
		// Getting data from POST request body
		decoder := json.NewDecoder(req.Body)

		// Structure for the PUT request json body
		type test_struct2 struct {
			X int
			Y int
		}

		var t test_struct2

		// Decoding JSON into test struct
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}

		prod := t.X * t.Y

		//storing into struct
		a := Solution{Result: prod}
		aBytes, err := json.Marshal(a)

		log.Println(err)
		log.Println(string(aBytes))

		fmt.Fprintf(w, string(aBytes))

		//perform multiplication
		//fmt.Fprintf(w, "Answer = %v", prod)

	} else if req.Method == "POST" {
		//check for POST req
		// Getting data from POST request body
		decoder := json.NewDecoder(req.Body)

		// Structure for the POST request json body
		type test_struct1 struct {
			X int
			Y int
		}

		var t test_struct1

		// Decoding JSON into test struct
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}

		result := 0

		if t.Y == 0 {
			fmt.Fprintf(w, "Can't perform division")
			return
		} else {
			result = t.X / t.Y
		}

		//storing into struct
		a := Solution{Result: result}
		aBytes, err := json.Marshal(a)

		log.Println(err)
		log.Println(string(aBytes))

		fmt.Fprintf(w, string(aBytes))

		//perform division
		//fmt.Fprintf(w, "Answer = %v", result)

	} else {

		//display message
		fmt.Fprint(w, "Babooshka")
	}

}
