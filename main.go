package main

import (
	"fmt"
	"os"
)

type DocConfig struct {
	// 1. mysql 2. oracle 3. mssql
	dbType   int
	host     string
	port     int
	user     string
	password string
	database string
	sid      string
	// 1. docsify
	docType int
}

var config = DocConfig{}

func main() {
	fmt.Println("choose database:\n1:MySQL\n2:Oracle\n3:SQL Server\n" +
		"Select the appropriate numbers choose database type\n" +
		"(Enter 'ctrl + c' to cancel):\n ")
	// db type
	fmt.Scanln(&config.dbType)
	if config.dbType < 1 || config.dbType > 4 {
		fmt.Println("wrong number, will exit ...")
		os.Exit(0)
	}
	GetDefaultConfig()
	// db host
	fmt.Println("input host (default 127.0.0.1) :")
	fmt.Scanln(&config.host)
	// db port
	fmt.Printf("input port (default %d) :\n", config.port)
	fmt.Scanln(&config.port)
	// db user
	fmt.Printf("input username (default %s) :\n", config.user)
	fmt.Scanln(&config.user)
	// db password
	fmt.Println("input password (default 123456) :")
	fmt.Scanln(&config.password)
	// db name
	if config.dbType == 2 {
		fmt.Println("input sid:")
		fmt.Scanln(&config.sid)
	} else {
		fmt.Println("input database name:")
		fmt.Scanln(&config.database)
	}
	// doc type
	fmt.Println("input doc type (default docsify) :")
	fmt.Scanln(&config.docType)
	// generate
	NewGenerate(&config)
}

func GetDefaultConfig() {
	if config.dbType == 1 {
		config.port = 3306
		config.user = "root"
	}
	if config.dbType == 2 {
		config.port = 1521
		config.user = ""
	}
	if config.dbType == 3 {
		config.port = 1433
		config.user = "sa"
	}
}
