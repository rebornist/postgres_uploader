package uploadData

import (
	"context"
	"fmt"
	"postgresUploader/configs"
)

func Upload(tn, fn string) {
	// 엑셀 데이터 추출
	op := TableDataReader(fn)

	// 추출된 엑셀 데이터를 쿼리로 변환
	qs := CreateQuery(tn, op)

	// DB 접속
	db := configs.ConnectDb()

	// begin a transaction
	tx := db.Begin()

	for _, q := range qs {
		if err := tx.Exec(q).Error; err != nil {
			panic(err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	fmt.Printf("%s Table %d's data upload Success\n", tn, len(qs))
}

func MongoDBUpload(tn, fn string) {

	// DB 접속
	client := configs.ConnectMongo()

	col := configs.GetDefaultCollection(client)

	// 엑셀 데이터 추출
	op := TableDataReader(fn)

	doc := CreateDocument(tn, op)

	for _, d := range doc {
		if _, err := col.InsertOne(context.TODO(), d); err != nil {
			panic(err)
		}
	}

	client.Disconnect(context.TODO())

}
