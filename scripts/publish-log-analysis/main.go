package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type collectionJSON struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	PublishDate      string `json:"publishDate"`
	ReleaseURI       string `json:"releaseUri"`
	PublishStartDate string `json:"publishStartDate"`
	PublishEndDate   string `json:"publishEndDate"`

	FileCountByType map[string]int
	WordCountByType map[string]int
	ByteCountByType map[string]int
}

type dataJSON struct {
	Type string `json:"type"`

	Sections []dataJSONMarkdownSection `json:"sections"`
}

type dataJSONMarkdownSection struct {
	Markdown string `json:"markdown"`
}

var collectionData = make(map[string]*collectionJSON)

// const rootDir = "/home/ec2-user/publish-log-test"

const rootDir = "/publish-archive"

func main() {
	// 1. Get list of collections
	// 2. Read collection data
	// 3. Find data.json inside collection
	// 4. Aggregate

	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		panic(err)
	}

	var fileTypes = make(map[string]int)

	lg := len(files)
	log.Printf("Found %d files to process", lg)

	for i, f := range files {
		if f.IsDir() {
			jsonFile := rootDir + "/" + f.Name() + ".json"
			b, err := ioutil.ReadFile(jsonFile)
			if err != nil {
				log.Printf("collection: [%s] cannot read file: [%s]", f.Name, jsonFile)
				continue
			}

			var c collectionJSON
			err = json.Unmarshal(b, &c)
			if err != nil {
				log.Printf("collection: [%s] invalid json: [%s]", f.Name, jsonFile)
				continue
			}

			c.FileCountByType = make(map[string]int)
			c.WordCountByType = make(map[string]int)
			c.ByteCountByType = make(map[string]int)

			collectionData[c.ID] = &c

			err = filepath.Walk(rootDir+"/"+f.Name(), func(path string, info os.FileInfo, err error) error {
				if err != nil {
					log.Printf("walk start failure: [%s]", f.Name)
					return err
				}
				if info.Name() == "data.json" {
					// log.Println("Found data.json:", path)
					b, err = ioutil.ReadFile(path)
					if err != nil {
						log.Printf("collection: [%s] cannot read data.json file: [%s]", f.Name, path)
						return err
					}

					var d dataJSON
					err = json.Unmarshal(b, &d)
					if err != nil {
						// log.Printf("Invalid data.json at %s: %s", path, err)
					}

					if _, ok := c.FileCountByType[d.Type]; !ok {
						c.FileCountByType[d.Type] = 1
					} else {
						c.FileCountByType[d.Type]++
					}
					if _, ok := fileTypes[d.Type]; !ok {
						fileTypes[d.Type] = 1
					}

					if _, ok := c.WordCountByType[d.Type]; !ok {
						c.WordCountByType[d.Type] = 0
					}
					if _, ok := c.ByteCountByType[d.Type]; !ok {
						c.ByteCountByType[d.Type] = 0
					}

					for _, s := range d.Sections {
						c.WordCountByType[d.Type] += len(strings.Split(s.Markdown, " "))
						c.ByteCountByType[d.Type] += len(s.Markdown)
					}
				}
				return nil
			})
			if err != nil {
				log.Printf("returned walk failure: [%s] error: [%s] ", f.Name, err.Error)
				continue
			}
		}

		if i%500 == 0 {
			log.Printf("Processed: %d/%d", i, lg)
		}
	}

	var typeNames []string
	for k := range fileTypes {
		typeNames = append(typeNames, k)
	}

	headers := []string{"ID", "Name", "ReleaseURI", "Type", "PublishDate", "PublishStartDate", "PublishEndDate"}
	for _, t := range typeNames {
		headers = append(headers, "Type_"+t+"_Count")
		headers = append(headers, "Type_"+t+"_Words")
		headers = append(headers, "Type_"+t+"_Bytes")
	}

	csvw := csv.NewWriter(os.Stdout)
	csvw.Write(headers)

	log.Printf("Found %d collections", len(collectionData))
	count := 0
	for _, v := range collectionData {
		row := []string{v.ID, v.Name, v.ReleaseURI, v.Type, v.PublishDate, v.PublishStartDate, v.PublishEndDate}
		for _, t := range typeNames {
			if v2, ok := v.FileCountByType[t]; ok {
				row = append(row, fmt.Sprintf("%d", v2))
			} else {
				row = append(row, "0")
			}
			if v2, ok := v.WordCountByType[t]; ok {
				row = append(row, fmt.Sprintf("%d", v2))
			} else {
				row = append(row, "0")
			}
			if v2, ok := v.ByteCountByType[t]; ok {
				row = append(row, fmt.Sprintf("%d", v2))
			} else {
				row = append(row, "0")
			}
		}
		csvw.Write(row)

		count++
		if count%50 == 0 {
			csvw.Flush()
		}
	}
	csvw.Flush()
}
