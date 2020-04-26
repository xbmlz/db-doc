package main

import (
	"db-doc/database"
	"db-doc/model"
	"fmt"
	"os"
)

var dbConfig model.DbConfig

func main() {
	fmt.Println("choose database:\n1:MySQL\n2:SQL Server\n" +
		"Select the appropriate numbers choose database type\n" +
		"(Enter 'ctrl + c' to cancel): ")
	// db type
	fmt.Scanln(&dbConfig.DbType)
	if dbConfig.DbType < 1 || dbConfig.DbType > 2 {
		fmt.Println("wrong number, will exit ...")
		os.Exit(0)
	}
	GetDefaultConfig()
	// db host
	fmt.Println("input host (default 127.0.0.1) :")
	fmt.Scanln(&dbConfig.Host)
	// db port
	fmt.Printf("input port (default %d) :\n", dbConfig.Port)
	fmt.Scanln(&dbConfig.Port)
	// db user
	fmt.Printf("input username (default %s) :\n", dbConfig.User)
	fmt.Scanln(&dbConfig.User)
	// db password
	fmt.Println("input password (default 123456) :")
	fmt.Scanln(&dbConfig.Password)
	// db name
	if dbConfig.DbType == 2 {
		fmt.Println("input sid:")
		fmt.Scanln(&dbConfig.Sid)
	} else {
		fmt.Println("input database name:")
		fmt.Scanln(&dbConfig.Database)
	}
	// doc type
	fmt.Println("choose doc type (default Docsify) :\n1:Docsify\n2:Gitbook")
	fmt.Scanln(&dbConfig.DocType)
	// generate
	database.Generate(&dbConfig)
}

// GetDefaultConfig get default config
func GetDefaultConfig() {
	dbConfig.Host = "127.0.0.1"
	dbConfig.Password = "123456"
	dbConfig.DocType = 1
	if dbConfig.DbType == 1 {
		dbConfig.Port = 3306
		dbConfig.User = "root"
	}
	if dbConfig.DbType == 2 {
		dbConfig.Port = 1433
		dbConfig.User = "sa"
	}
}
