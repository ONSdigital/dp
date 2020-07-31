package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

type collectionJSON struct {
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	Type             string     `json:"type"`
	PublishDate      string     `json:"publishDate"`
	PublishStartDate *time.Time `json:"publishStartDate"`
	PublishEndDate   *time.Time `json:"publishEndDate"`

	Duration  int64
	FileCount int64
	FileSize  int64
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

			collectionData[c.ID] = &c

			if c.PublishStartDate != nil && c.PublishEndDate != nil {
				c.Duration = int64(c.PublishEndDate.Sub(*c.PublishStartDate)) / int64(time.Millisecond)
			}

			err = filepath.Walk(rootDir+"/"+f.Name(), func(path string, info os.FileInfo, err error) error {
				if err != nil {
					log.Printf("walk start failure: [%s]", f.Name)
					return err
				}
				c.FileCount++
				c.FileSize += info.Size()

				return nil
			})
		}

		if i%500 == 0 {
			log.Printf("Processed: %d/%d", i, lg)
		}
	}

	headers := []string{"ID", "Name", "Type", "PublishDate", "PublishStartDate", "PublishEndDate", "Duration", "FileCount", "FileSize"}

	csvw := csv.NewWriter(os.Stdout)
	csvw.Write(headers)

	log.Printf("Found %d collections", len(collectionData))
	count := 1
	for _, v := range collectionData {
		sDate, eDate := "", ""
		if v.PublishStartDate != nil {
			sDate = v.PublishStartDate.String()
		}
		if v.PublishEndDate != nil {
			eDate = v.PublishEndDate.String()
		}
		row := []string{v.ID, v.Name, v.Type, v.PublishDate, sDate, eDate, fmt.Sprintf("%d", v.Duration), fmt.Sprintf("%d", v.FileCount), fmt.Sprintf("%d", v.FileSize)}
		csvw.Write(row)

		count++
		if count%50 == 0 {
			csvw.Flush()
		}
	}
	csvw.Flush()
}
