package configs

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB 접속 정보
type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func ConnectDb() *gorm.DB {

	// 웹 서비스 정보 중 데이터베이스 정보 추출
	var DB Database
	getInfo, err := GetServiceInfo("database")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(getInfo, &DB)

	// gorm DB 접속
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul", DB.Host, DB.User, DB.Password, DB.Name, DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
