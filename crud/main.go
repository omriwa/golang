package main

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Name string
}

func getAllPersons(channel chan string) {
	db, dbError := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/samp")

	defer db.Close()

	if dbError != nil {
		panic(dbError.Error())
	}

	result, queryError := db.Query("SELECT * FROM Persons")

	if queryError != nil {
		channel <- "QUERY ERROR"
	}

	defer result.Close()

	output := ""
	p := Person{}

	for result.Next() {

		result.Scan(&p.Name)

		output += p.Name + "\n"
	}

	channel <- output
}

func insertName(name string, channel chan int) {
	db, dbError := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/samp")

	defer db.Close()

	if dbError != nil {
		panic(dbError.Error())
	}

	_, queryError := db.Query(fmt.Sprintf("INSERT INTO Persons(Name) VALUES (\"%v\")", name))

	if queryError != nil {
		channel <- 500
	}

	channel <- 200
}

func main() {
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			return
		}
		bytedata, err := ioutil.ReadAll(r.Body)

		if err != nil {
			io.WriteString(w, "500")

		}

		c := make(chan int)
		reqBodyString := string(bytedata)

		go insertName(reqBodyString, c)

		io.WriteString(w, string(<-c))
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			return
		}
		c := make(chan string)

		go getAllPersons(c)

		io.WriteString(w, <-c)
	})

	http.ListenAndServe(":3000", nil)
}
