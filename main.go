package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path,has_env := os.LookupEnv("DB_URL")
	if !has_env {
		log.Fatal("No DB_URL env supplied")
	}
	fmt.Println(path)
}