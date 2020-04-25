package database

import (
	"database/sql"
	"db-doc/model"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

var dbConfig model.DbConfig

// Generate generate doc
func Generate(config *model.DbConfig) {
	dbConfig = *config
	fmt.Println(dbConfig)
	db := initDB()
	if db == nil {
		fmt.Println("init databse err")
		os.Exit(1)
	}
	defer db.Close()
	tables := getTableInfo(db)
	fmt.Println(tables)
}

// InitDB 初始化数据库
func initDB() *sql.DB {
	var (
		dbURL  string
		dbType string
	)
	if dbConfig.DbType == 1 {
		// https://github.com/go-sql-driver/mysql/
		dbType = "mysql"
		// <username>:<password>@<host>:<port>/<database>
		dbURL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	}
	if dbConfig.DbType == 2 {
		// TODO
	}
	if dbConfig.DbType == 3 {
		// https://github.com/denisenkom/go-mssqldb
		dbType = "mssql"
		// server=%s;database=%s;user id=%s;password=%s;port=%d;encrypt=disable
		dbURL = fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d;encrypt=disable",
			dbConfig.Host, dbConfig.Database, dbConfig.User, dbConfig.Password, dbConfig.Port)
	}
	fmt.Println(dbURL)
	db, err := sql.Open(dbType, dbURL)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return db
}

// getTableInfo 获取表信息
func getTableInfo(db *sql.DB) []model.Table {
	// find all tables
	tables := make([]model.Table, 0)
	rows, err := db.Query(getTableSQL())
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var table model.Table
	for rows.Next() {
		rows.Scan(&table.TableName, &table.TableComment)
		tables = append(tables, table)
	}
	for i := range tables {
		columns := getColumnInfo(db, tables[i].TableName)
		tables[i].ColList = columns
	}
	return tables
}

// getColumnInfo 获取列信息
func getColumnInfo(db *sql.DB, tableName string) []model.Column {
	columns := make([]model.Column, 0)
	rows, err := db.Query(getColumnSQL(tableName))
	if err != nil {
		fmt.Println(err)
	}
	var column model.Column
	for rows.Next() {
		rows.Scan(&column.ColName, &column.ColType, &column.ColKey, &column.IsNullable, &column.ColComment)
		columns = append(columns, column)
	}
	return columns
}

// getTableSQL
func getTableSQL() string {
	var sql string
	if dbConfig.DbType == 1 {
		sql = fmt.Sprintf("select table_name as TableName, table_comment as TableComment from information_schema.tables where table_schema = '%s'",
			dbConfig.Database)
	}
	if dbConfig.DbType == 2 {
		// TODO
	}
	if dbConfig.DbType == 3 {
		sql = fmt.Sprintf(`
		select * from (
			select cast(so.name as varchar(500)) as TableName, 
			cast(sep.value as varchar(500)) as TableComment
			from sysobjects so
			left JOIN sys.extended_properties sep on sep.major_id=so.id and sep.minor_id=0
			where (xtype='U' or xtype='v')
		) t 
		`)
	}
	return sql
}

// getColumnSQL
func getColumnSQL(tableName string) string {
	var sql string
	if dbConfig.DbType == 1 {
		sql = fmt.Sprintf("select column_name as ColName, column_type as ColType, column_key as ColKey, is_nullable as IsNullable, column_comment as ColComment"+
			" from information_schema.columns where table_schema = '%s' and table_name = '%s'",
			dbConfig.Database, tableName)
	}
	if dbConfig.DbType == 2 {
		// TODO
	}
	if dbConfig.DbType == 3 {
		sql = fmt.Sprintf(`
		SELECT ColName = C.name, 
			   ColKey = ISNULL(IDX.PrimaryKey, NULL), 
			   ColType = T.name, 
			   IsNullable = CASE WHEN C.is_nullable = 1 THEN N'是' ELSE N'否' END, 
			   ColComment = ISNULL(CAST(PFD.[value] AS VARCHAR(500)), NULL)
		FROM sys.columns C
		INNER JOIN sys.objects O ON C.object_id = O.object_id AND O.type = 'U' AND O.is_ms_shipped = 0
		INNER JOIN sys.types T ON C.user_type_id = T.user_type_id
		LEFT  JOIN sys.default_constraints D ON C.object_id = D.parent_object_id AND C.column_id = D.parent_column_id AND C.default_object_id = D.[object_id]
		LEFT  JOIN sys.extended_properties PFD ON PFD.class = 1 AND C.[object_id] = PFD.major_id AND C.column_id = PFD.minor_id
		LEFT  JOIN sys.extended_properties PTB ON PTB.class = 1 AND PTB.minor_id = 0 AND C.[object_id] = PTB.major_id
		LEFT  JOIN 
			(
				SELECT IDXC.[object_id], 
					IDXC.column_id, 
					Sort = CASE INDEXKEY_PROPERTY(IDXC.[object_id], IDXC.index_id, IDXC.index_column_id, 'IsDescending')
						WHEN 1 THEN 'DESC'
						WHEN 0 THEN 'ASC'
						ELSE ''
						END, 
					PrimaryKey = CASE WHEN IDX.is_primary_key = 1 THEN N'PRI' ELSE NULL END, 
					IndexName = IDX.Name
				FROM sys.indexes IDX
				INNER JOIN sys.index_columns IDXC ON IDX.[object_id] = IDXC.[object_id] AND IDX.index_id = IDXC.index_id
				LEFT  JOIN sys.key_constraints KC ON IDX.[object_id] = KC.[parent_object_id] AND IDX.index_id = KC.unique_index_id
				INNER JOIN (
					SELECT [object_id], Column_id, index_id = MIN(index_id)
					FROM sys.index_columns
					GROUP BY [object_id], Column_id
				) IDXCUQ ON IDXC.[object_id] = IDXCUQ.[object_id] AND IDXC.Column_id = IDXCUQ.Column_id AND IDXC.index_id = IDXCUQ.index_id
			) IDX ON C.[object_id] = IDX.[object_id] AND C.column_id = IDX.column_id
		WHERE O.name = '%s'
		`, tableName)
	}
	return sql
}
