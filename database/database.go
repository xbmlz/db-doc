package database

import (
	"database/sql"
	"db-doc/doc"
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
	db := initDB()
	if db == nil {
		fmt.Println("init database err")
		os.Exit(1)
	}
	defer db.Close()
	tables := getTableInfo(db)
	// create
	doc.CreateDoc(dbConfig.Database, tables)
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
		// https://github.com/denisenkom/go-mssqldb
		dbType = "mssql"
		// server=%s;database=%s;user id=%s;password=%s;port=%d;encrypt=disable
		dbURL = fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d;encrypt=disable",
			dbConfig.Host, dbConfig.Database, dbConfig.User, dbConfig.Password, dbConfig.Port)
	}
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
		if len(table.TableComment) == 0 {
			table.TableComment = table.TableName
		}
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
		rows.Scan(&column.ColName, &column.ColType, &column.ColKey, &column.IsNullable, &column.ColComment, &column.ColDefault)
		columns = append(columns, column)
	}
	return columns
}

// getTableSQL
func getTableSQL() string {
	var sql string
	if dbConfig.DbType == 1 {
		sql = fmt.Sprintf(`
			select table_name    as TableName, 
			       table_comment as TableComment
			from information_schema.tables 
			where table_schema = '%s'
		`, dbConfig.Database)
	}
	if dbConfig.DbType == 2 {
		sql = fmt.Sprintf(`
		select * from (
			select cast(so.name as varchar(500)) as TableName, 
			cast(sep.value as varchar(500))      as TableComment
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
		sql = fmt.Sprintf(`
			select column_name as ColName,
			column_type        as ColType,
			column_key         as ColKey,
			is_nullable        as IsNullable,
			column_comment     as ColComment,
			column_default     as ColDefault
			from information_schema.columns 
			where table_schema = '%s' and table_name = '%s'
		`, dbConfig.Database, tableName)
	}
	if dbConfig.DbType == 2 {
		sql = fmt.Sprintf(`
		SELECT 
			ColName = a.name,
			ColType = b.name + '(' + cast(COLUMNPROPERTY(a.id, a.name, 'PRECISION') as varchar) + ')',
			ColKey  = case when exists(SELECT 1
										FROM sysobjects
										where xtype = 'PK'
										and name in (
											SELECT name
											FROM sysindexes
											WHERE indid in (
												SELECT indid
												FROM sysindexkeys
												WHERE id = a.id AND colid = a.colid
										))) then 'PRI'
							else '' end,
			IsNullable = case when a.isnullable = 1 then 'YES' else 'NO' end,
			ColComment = isnull(g.[value], ''),
			ColDefault = isnull(e.text, '')
		FROM syscolumns a
				left join systypes b on a.xusertype = b.xusertype
				inner join sysobjects d on a.id = d.id and d.xtype = 'U' and d.name <> 'dtproperties'
				left join syscomments e on a.cdefault = e.id
				left join sys.extended_properties g on a.id = g.major_id and a.colid = g.minor_id
				left join sys.extended_properties f on d.id = f.major_id and f.minor_id = 0
		where d.name = '%s'
		order by a.id, a.colorder
		`, tableName)
	}
	return sql
}
