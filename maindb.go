package main

//Using a docker container of mysql setting up employee db and
//make an API to create an employee.

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/employee", employeeHandler)

	http.ListenAndServe(":8000", nil)
}

func employeeHandler(w http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("mysql", "root:sam123@tcp(localhost:2000)/test1")
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	//Creating employee struct

	type employee struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		City string `json:"city"`
	}

	emp := employee{}

	if req.Method == "GET" {

		id := req.URL.Query().Get("id")

		log.Println(id)
		i, _ := strconv.Atoi(id)

		row := db.QueryRow("Select * FROM employee where id=?", i)

		//error in row

		err = row.Scan(&emp.Id, &emp.Name, &emp.City)

		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Crash!")
			return
		}

		empBytes, err := json.Marshal(emp)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Crash!!!")
			return
		}

		fmt.Fprintf(w, string(empBytes))

	} else if req.Method == "POST" {

		// Getting data from POST request body
		decoder := json.NewDecoder(req.Body)

		type body_struct struct {
			Name string
			City string
		}

		var one body_struct

		err := decoder.Decode(&one)
		if err != nil {
			log.Println(err)
			return
		}

		//Adding row
		rowadd, err := db.Exec("Insert into employee (name,city) VALUES (?,?)", one.Name, one.City)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "error in addition in table!")
			return
		}

		output, err := rowadd.LastInsertId()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Conversion")
			return
		}
		fmt.Fprintf(w, "Added = %v ", output)

	}

	//error in row

	// err = row.Scan(&emp.Id, &emp.Name, &emp.City)

	// if err != nil {
	// 	log.Println(err)
	// 	fmt.Fprintf(w, "Crash!")
	// 	return
	// }

	// empBytes, err := json.Marshal(emp)
	// if err != nil {
	// 	log.Println(err)
	// 	fmt.Fprintf(w, "Crash!!!")
	// 	return
	// }

	// fmt.Fprintf(w, string(empBytes))

}
