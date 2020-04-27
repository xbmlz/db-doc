package doc

import (
	"db-doc/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

const docsifyHTML = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Databse Document</title>
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
	<meta name="description" content="Description">
	<meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
	<link rel="stylesheet" href="//unpkg.com/docsify/lib/themes/vue.css">
	</head>
	<body>
	<div data-app id="main">加载中</div>
	<script>
		window.$docsify = {
			el: '#main',
			name: '',
			repo: '',
			search: 'auto',
			loadSidebar: true
		}
	</script>
	<script src="//unpkg.com/docsify/lib/docsify.min.js"></script>
	<script src="//unpkg.com/docsify/lib/plugins/search.js"></script>
	</body>
	</html>
`

// CreateDoc create doc
func CreateDoc(docType int, dbName string, tables []model.Table) {
	dir, _ := os.Getwd()
	docPath := path.Join(dir, dbName)
	createDir(docPath)
	if docType == 1 {
		createDocsify(docPath, dbName, tables)
	}
}

// createDocsify create _siderbar.md
func createDocsify(docPath string, dbName string, tables []model.Table) {
	var sidebar []string
	var readme []string
	sidebar = append(sidebar, "* [数据库文档](README.md)")
	for i := range tables {
		readme = append(readme, fmt.Sprintf("# %s数据库文档", dbName))
		readme = append(readme, fmt.Sprintf("- [%s](%s.md)", tables[i].TableComment, tables[i].TableName))
		sidebar = append(sidebar, fmt.Sprintf("* [%s](%s.md)", tables[i].TableComment, tables[i].TableName))
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
	// create readme.md
	readmeStr := strings.Join(sidebar, "\r\n")
	writeToFile(path.Join(docPath, "README.md"), readmeStr)
	// create _sidebar.md
	sidebarStr := strings.Join(sidebar, "\r\n")
	writeToFile(path.Join(docPath, "_sidebar.md"), sidebarStr)
	// create index.html
	writeToFile(path.Join(docPath, "index.html"), docsifyHTML)
	// create .nojekyll
	writeToFile(path.Join(docPath, ".nojekyll"), "")
	fmt.Println("doc generate successfully!")
	// run server
	runServer(docPath)
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
func runServer(dir string) {
	http.Handle("/", http.FileServer(http.Dir(dir)))
	fmt.Println("doc server is runing : http://127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
