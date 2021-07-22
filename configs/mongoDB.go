package configs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (client *mongo.Client) {

	// 웹 서비스 정보 중 데이터베이스 정보 추출
	var DB Database
	getInfo, err := GetServiceInfo("mongodb")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(getInfo, &DB)

	// Auth에러 처리를 위한 client option 구성
	uri := fmt.Sprintf("mongodb://%s:%s", DB.Host, DB.Port)
	clientOptions := options.Client().ApplyURI(uri)

	// MongoDB 연결
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connection Made")
	return client
}

func GetDefaultCollection(client *mongo.Client) (collection *mongo.Collection) {
	return client.Database("prs_search").Collection("search")
}
