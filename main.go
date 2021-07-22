package main

import (
	"os"
	"postgresUploader/controllers/mkTable"
	"postgresUploader/controllers/uploadData"
)

func main() {
	od := os.Args[1]
	tn := os.Args[2]
	fn := os.Args[3]
	m := ""
	if len(os.Args) > 4 {
		m = os.Args[4]
	}

	switch od {
	case "--mt":
		mkTable.CreateTable(tn, fn)
	case "--upload":
		if m == "mongodb" {
			uploadData.MongoDBUpload(tn, fn)
		} else {
			uploadData.Upload(tn, fn)
		}
	}

}
