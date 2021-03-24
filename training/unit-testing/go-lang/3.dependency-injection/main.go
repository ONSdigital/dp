package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/ONSdigital/dp/training/unitTesting/goLang/3.dependencyInjection/datastore"
	"github.com/ONSdigital/dp/training/unitTesting/goLang/3.dependencyInjection/service"
)

// This is a training application made by Office to demonstrate testing in Go
func main() {
	// Define connection string for MongoDB
	conStr := "dbname=corporatePostgres sslmode=enabled"
	db, _ := sql.Open("postgres", conStr)
	// Create a data store dependency with the db connection
	store := datastore.NewDatastore(db)
	// Create the service by injecting the store as a dependency
	service := &service.Service{Store: store}
	// The following code implements a simple command line app to read the ID as input
	// and output the validity of the result of the entry with that ID in the database
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ID, _ := strconv.Atoi(scanner.Text())
		err := service.GetNumber(ID)
		if err != nil {
			fmt.Printf("result invalid: %v", err)
			continue
		}
		fmt.Println("result valid")
	}
}
