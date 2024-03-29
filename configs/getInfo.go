package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 웹 서비스 정보 받아오기
func getServiceInfoInit() map[string]interface{} {
	var info map[string]interface{}
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	data, err := os.Open(fmt.Sprintf("%s/configs/%s", pwd, "config.json"))
	if err != nil {
		fmt.Println(err)
	}
	byteData, err := ioutil.ReadAll(data)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(byteData, &info)
	return info
}

func GetServiceInfo(name string) ([]byte, error) {
	// 웹 서비스 정보 중 데이터베이스 정보 추출
	getInfo, err := json.Marshal(getServiceInfoInit()[name])
	if err != nil {
		fmt.Println(err)
	}
	return getInfo, err
}
