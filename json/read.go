package json

import (
	"404Chat/controller"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadFileJson(fileName string) error {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteJson, err := ioutil.ReadAll(jsonFile)

	var users []controller.Account

	_ = json.Unmarshal(byteJson, &users)
	users = append(users, user)
	for i := 0; i < len(users); i++ {
		fmt.Println("username" + users[i].UserName)
		fmt.Println("password" + users[i].Password)
	}
	return jsonFile.Sync()
}

func ReadJson(fileName string, filter func(map[string]interface{}) bool) []map[string]interface{} {
	var data []map[string]interface{}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(file, &data)

	var fileData []map[string]interface{}

	for _, dt := range data {
		if filter(dt) {
			fileData = append(fileData, dt)
		}
	}
	return fileData
}

func ReadJsonToken(fileName string, filter func(map[string]interface{}) bool) []map[string]interface{} {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	var fileData []map[string]interface{}

	_, _ = decoder.Token()

	var data map[string]interface{}

	for decoder.More() {
		_ = decoder.Decode(&data)

		if filter(data) {
			fileData = append(fileData, data)
		}
	}
	return fileData
}
