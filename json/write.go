package json

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"os"
)

func WriteJson(fileName string) error{
	file,err:=os.OpenFile(fileName,os.O_CREATE,os.ModePerm)
	if err!=nil{
		log.Fatal(err)
	}
	defer file.Close()

	encoder:=json.NewEncoder(file)
	JsonString := encoder.Encode(10000)
	err = ioutil.WriteFile(fileName, JsonString, os.ModePerm)
	if err!=nil{
		log.Fatal(err)
	}
	return file.Sync()
}