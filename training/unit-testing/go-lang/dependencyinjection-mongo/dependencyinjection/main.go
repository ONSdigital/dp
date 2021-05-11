package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"github.com/ONSdigital/dp/training/unit-testing/go-lang/dependencyinjection/datastore"
	service "github.com/ONSdigital/dp/training/unit-testing/go-lang/dependencyinjection/service"
	"github.com/benweissmann/memongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strconv"
)
// main is a simple implementation of a database routine which takes in a value from the
// user and get's the result from the database.
func main() {
	//Spinning up in memory mongo database
	mongoServer, err := memongo.Start("4.0.5")
	if (err != nil) {
		panic (err)
	}
	defer mongoServer.Stop()
	fmt.Printf("Connection String:" + mongoServer.URI()+"\n");

	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//var collection *mongo.Collection = client.Database("tasker").Collection("tasks")
	var collection = client.Database("tasker").Collection("tasks")




	// Check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	connectAndDoStuff(mongoServer.URI(), memongo.RandomDatabase())
	
	conStr := mongoServer.URI()
	d, err := sql.Open("", conStr)
	if err != nil {
		fmt.Printf("result invalid: %v", err)
	}
	// Create a datastore using the connection string
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

func connectAndDoStuff(uri string, database string) {
	
}
