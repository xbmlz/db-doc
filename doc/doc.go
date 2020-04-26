package doc

import (
	"db-doc/model"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// CreateDoc create doc
func CreateDoc(docType int, dbName string, tables []model.Table) {
	dir, _ := os.Getwd()
	docPath := path.Join(dir, dbName)
	createDir(docPath)
	if docType == 1 {
		createDocsify(docPath, tables)
	}
}

// createDocsify create _siderbar.md
func createDocsify(docPath string, tables []model.Table) {
	var siderbar []string
	siderbar = append(siderbar, "* [数据库文档](README.md)")
	for i := range tables {
		siderbar = append(siderbar, fmt.Sprintf("[%s](%s.md)", tables[i].TableComment, tables[i].TableName))
		var tableMd []string
		tableMd = append(tableMd, fmt.Sprintf("# %s(%s)", tables[i].TableComment, tables[i].TableName))
		tableMd = append(tableMd, "| 列名 | 类型 | KEY | 可否为空 | 默认值 | 注释 |")
		tableMd = append(tableMd, "| ---- | ---- | ---- | ---- | ---- | ----  |")
		// create table.md
		cols := tables[i].ColList
		for j := range cols {
			tableMd = append(tableMd, fmt.Sprintf("| %s | %s | %s | %s | %s | %s |",
				cols[j].ColName, cols[j].ColType, cols[j].ColKey, cols[j].IsNullable, "", cols[j].ColComment))
		}
		tableStr := strings.Join(tableMd, "\r\n")
		writeToFile(path.Join(docPath, tables[i].TableName+".md"), tableStr)
	}
	// create _siderbar.md
	siderbarStr := strings.Join(siderbar, "\r\n")
	writeToFile(path.Join(docPath, "_siderbar.md"), siderbarStr)
}

// createDir
func createDir(dirPath string) error {
	if !isExist(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		return err
	}
	return nil
}

// isExist
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// writeToFile write file
func writeToFile(path, content string) {
	if err := ioutil.WriteFile(path, []byte(content), 777); err != nil {
		os.Exit(1)
		log.Println(err.Error())
	}
}

// runServer run http static server
func runServer() {

}
