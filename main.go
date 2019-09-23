package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var DB *sql.DB

func main() {
	path,has_env := os.LookupEnv("DB_URL")
	if !has_env {
		log.Fatal("No DB_URL env supplied")
	}
	var err error
	DB, err = sql.Open("postgres",path)
	if err != nil {
		log.Fatal(err)
	}
	file, ioErr := ioutil.ReadFile("./lab-test.sql")
	if ioErr != nil {
		log.Fatal("Read Fail")
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		result, err := DB.Exec(request)
		if err != nil {
			log.Fatal("DataBase Error")
		}
		fmt.Println(result)
	}
}