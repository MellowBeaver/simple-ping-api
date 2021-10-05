package main

//Using a docker container of mysql setting up employee db and
//make an API to create an employee.

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/employee", employeeHandler)

	http.ListenAndServe(":8000", nil)
}

func employeeHandler(w http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("mysql", "root:password@tcp(localhost:2000)/test1")
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	id := req.URL.Query().Get("id")

	log.Println(id)

	row := db.QueryRow("Select * FROM test1.employee where id=?", id)

	//Adding row

	name := req.URL.Query().Get("name")
	city := req.URL.Query().Get("city")

	rowadd := db.QueryRow("Insert into employee (id, name, city) VALUES (?, ?, ?);", id, name, city)

	//Creating employee struct

	type employee struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		City string `json:"city"`
	}

	emp := employee{}

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

}
