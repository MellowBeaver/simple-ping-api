package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var path = "/apiCaller"

func main() {

	http.HandleFunc("/apiCaller", pingHandler)

	http.ListenAndServe(":8000", nil)
}

func pingHandler(w http.ResponseWriter, req *http.Request) {

	//fmt.Fprintf(w, "API Caller\n")

	//take command using POST method
	if req.Method == "POST" {

		//conditions to perform functions
		input := req.URL.Query().Get("input")

		if input == "add" {
			//fmt.Fprintln(w, "Addition\n\n ")

			x := req.URL.Query().Get("x")
			y := req.URL.Query().Get("y")

			//call URL
			req, _ := http.NewRequest("GET", "http://localhost:8080/arithmetic?x="+x+"&y="+y, nil)

			//client stores response
			client := http.DefaultClient
			resp, err := client.Do(req)

			//log error
			if err != nil {
				log.Println(err)
			}

			//to check if other API responsive
			fmt.Println(resp.StatusCode)

			a, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(a))

			fmt.Fprintf(w, string(a))

		} else if input == "sub" {
			//fmt.Fprintln(w, "Subtraction\n\n ")

			//extracting from request body ****very imp****
			b, _ := ioutil.ReadAll(req.Body)

			req, _ := http.NewRequest("DELETE", "http://localhost:8080/arithmetic", bytes.NewBuffer(b))

			//client stores response
			client := http.DefaultClient
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
			}

			//to check if other API responsive
			fmt.Println(resp.StatusCode)

			a, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(a))
			fmt.Fprintln(w, string(a))

		} else if input == "mul" {
			//fmt.Fprintln(w, "Multiplication\n\n ")

			//extracting from request body ****very imp****
			b, _ := ioutil.ReadAll(req.Body)

			req, _ := http.NewRequest("PUT", "http://localhost:8080/arithmetic", bytes.NewBuffer(b))

			//client stores response
			client := http.DefaultClient
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
			}

			a, _ := ioutil.ReadAll(resp.Body)
			fmt.Fprintf(w, string(a))

		} else if input == "div" {
			//fmt.Fprintln(w, "Division\n\n ")

			//extracting from request body ****very imp****
			b, _ := ioutil.ReadAll(req.Body)

			req, _ := http.NewRequest("POST", "http://localhost:8080/arithmetic", bytes.NewBuffer(b))

			//client stores response
			client := http.DefaultClient
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
			}

			a, _ := ioutil.ReadAll(resp.Body)
			fmt.Fprintf(w, string(a))

		} else {
			fmt.Fprintln(w, "Enter value in input variable\n\nTry\nadd\nsub\nmul\ndiv")
		}
	} else {
		fmt.Fprintln(w, "Input via POST method")
	}

}
