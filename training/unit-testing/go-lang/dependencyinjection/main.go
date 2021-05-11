package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/ONSdigital/dp/training/unit-testing/go-lang/dependencyinjection/datastore"
	service "github.com/ONSdigital/dp/training/unit-testing/go-lang/dependencyinjection/service"
	"os"
	"strconv"
)

func main() {
	// Define connection string for MongoDB
	conStr := "dbname=aDatabase"
	d, _ := sql.Open("postgres", conStr)
	// Create a data store dependency with the d connection
	store := datastore.NewDatastore(d)
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
