package doc

import (
	"db-doc/model"
	"db-doc/util"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

const docsifyHTML = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Database Document</title>
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
	<script src="//unpkg.com/docsify/lib/plugins/search.min.js"></script>
	</body>
	</html>
`

// createOnlineDoc create _siderbar.md
func createOnlineDoc(docPath string, dbName string, tables []model.Table) {
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
		util.WriteToFile(path.Join(docPath, tables[i].TableName+".md"), tableStr)
	}
	// create readme.md
	readmeStr := strings.Join(sidebar, "\r\n")
	util.WriteToFile(path.Join(docPath, "README.md"), readmeStr)
	// create _sidebar.md
	sidebarStr := strings.Join(sidebar, "\r\n")
	util.WriteToFile(path.Join(docPath, "_sidebar.md"), sidebarStr)
	// create index.html
	util.WriteToFile(path.Join(docPath, "index.html"), docsifyHTML)
	// create .nojekyll
	util.WriteToFile(path.Join(docPath, ".nojekyll"), "")
	fmt.Println("doc generate successfully!")
	// run server
	runServer(docPath)
}

// runServer run http static server
func runServer(dir string) {
	http.Handle("/", http.FileServer(http.Dir(dir)))
	fmt.Println("doc server is runing : http://127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
