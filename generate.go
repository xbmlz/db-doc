package main

import (
	"database/sql"
	"fmt"
)

type Generate struct {
}

type Column struct {
	colName    string
	colType    string
	colKey     string
	isNullable string
	colComment string
}

type Table struct {
	tableName    string
	tableComment string
	colList      []Column
}

var dbConfig = DocConfig{}

func NewGenerate(config *DocConfig) {
	dbConfig = *config
	initDB()
}

// InitDB 初始化数据库
func initDB() *sql.DB {
	var (
		dbUrl  string
		dbType string
	)
	if dbConfig.dbType == 1 {
		// https://github.com/go-sql-driver/mysql/
		dbUrl = fmt.Sprintf("%s:%s@/%scharset=utf8", dbConfig.user, dbConfig.password, dbConfig.database)
	}
	db, err := sql.Open(dbType, dbUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	return db
}

// getUrl
func getUrl() {

}
