package main

import (
	"db-doc/database"
	"db-doc/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var dbConfig *model.DbConfig

func main() {
	Setup()
	// generate
	database.Generate(dbConfig)
}

// GetDefaultConfig get default config
func Setup() {
	GetConfig()
}


func GetConfig() *model.DbConfig {
	if dbConfig == nil {
		cfg, err := loadConfig()
		if err != nil {
			log.Fatal(err)
			return nil
		}
		dbConfig = &cfg
	}
	return dbConfig
}

func loadConfig() (config model.DbConfig, err error) {
	data, err := loadFile("./conf.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return
	}

	return
}

func loadFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(file)
}