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

	switch od {
	case "--mt":
		mkTable.CreateTable(tn, fn)
	case "--upload":
		uploadData.Upload(tn, fn)
	}

}
