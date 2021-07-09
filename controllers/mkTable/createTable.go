package mkTable

import (
	"fmt"
	"postgresUploader/configs"
)

func CreateTable(fn, tn string) {
	// 엑셀 데이터 추출
	op := TableDataReader(fn)

	// 추출된 엑셀 데이터를 쿼리로 변환
	q := CreateQuery(tn, op)

	// DB 접속
	db := configs.ConnectDb()

	if err := db.Exec(q).Error; err != nil {
		panic(err)
	}

	fmt.Printf("%s Table Create Success\n", tn)
}
