package main

import (
	"fmt"
	"os"
	"postgresUploader/configs"
	"postgresUploader/controllers/mkTable"
)

func main() {
	var q string
	db := configs.ConnectDb()

	od := os.Args[1]
	in := os.Args[2]
	fn := os.Args[3]

	switch od {
	case "--mt":
		mkTable.CreateTable(in, fn)
	case "--upload":
		fmt.Println("Upload datas")
	}

}
