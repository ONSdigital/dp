package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

type collectionJSON struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	PublishDate *time.Time `json:"publishDate"`
	ReleaseURI  string     `json:"releaseUri"`
}

var collectionData = make(map[string]*collectionJSON)

func main() {
	// 1. Get list of collections
	// 2. Read collection data
	// 3. Find data.json inside collection
	// 4. Aggregate

	files, err := ioutil.ReadDir("/home/ec2-user/publish-log")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if f.IsDir() {
			jsonFile := "/home/ec2-user/publish-log/" + f.Name() + ".json"
			b, err := ioutil.ReadFile(jsonFile)
			if err != nil {
				panic(err)
			}

			var c collectionJSON
			err = json.Unmarshal(b, &c)
			if err != nil {
				panic(err)
			}

			collectionData[c.ID] = &c

			err = filepath.Walk("/home/ec2-user/publish-log/"+f.Name(), func(path string, info os.FileInfo, err error) error {
				if err != nil {
					panic(err)
				}
				if info.Name() == "data.json" {
					log.Println("Found data.json:", path)
				}
				return nil
			})
			if err != nil {
				panic(err)
			}
		}
	}

	log.Printf("Found %d collections", len(collectionData))
}
