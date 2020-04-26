package doc

import (
	"db-doc/model"
	"fmt"
	"strings"
)

// saveFile save file
func saveFile() {

}

// runServer run http static server
func runServer() {

}

// createDocsify create _siderbar.md
func createDocsify(tables []model.Table) {
	var siderbar []string
	siderbar = append(siderbar, "* [数据库文档](README.md)")
	for i := range tables {
		siderbar = append(siderbar, fmt.Sprintf("[%s](%s.md)", tables[i].TableComment, tables[i].TableName))
		// TODO create table.md
	}
	// TODO create _siderbar.md
	strings.Join(siderbar, "\r\n")
}
