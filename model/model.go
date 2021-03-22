package model

// DbConfig 数据库配置
type DbConfig struct {
	DbType   int // 1. mysql  2. oracle 3. mssql
	DocType  int // 1. online 3. offline
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Sid      string
}

// Column info
type Column struct {
	ColName    string
	ColType    string
	ColKey     string
	IsNullable string
	ColComment string
	ColDefault string
}

// Table info
type Table struct {
	TableName    string
	TableComment string
	ColList      []Column
}
