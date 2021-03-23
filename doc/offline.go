package doc

import (
	"db-doc/model"
	"db-doc/util"
	"fmt"
	"github.com/russross/blackfriday"
	"path"
	"strings"
)

// createOfflineDoc create offline html、md、pdf、word
func createOfflineDoc(docPath string, dbName string, tables []model.Table) {
	var (
		docMdArr []string
		docMdStr string
	)
	// markdown
	for i := range tables {
		docMdArr = append(docMdArr, fmt.Sprintf("# %s(%s)", tables[i].TableComment, tables[i].TableName))
		docMdArr = append(docMdArr, "| 列名 | 类型 | KEY | 可否为空 | 默认值 | 注释 |")
		docMdArr = append(docMdArr, "| ---- | ---- | ---- | ---- | ---- | ----  |")
		// create table.md
		cols := tables[i].ColList
		for j := range cols {
			docMdArr = append(docMdArr, fmt.Sprintf("| %s | %s | %s | %s | %s | %s |",
				cols[j].ColName, cols[j].ColType, cols[j].ColKey, cols[j].IsNullable, "", cols[j].ColComment))
		}
		docMdArr = append(docMdArr, "")
	}
	docMdStr = strings.Join(docMdArr, "\r\n")
	util.WriteToFile(path.Join(docPath, dbName+".md"), docMdStr)
	// html
	docMdArr = append([]string{mdCss}, docMdArr...)
	docMdStr = strings.Join(docMdArr, "\r\n")
	fmt.Println("markdown generate successfully!")
	convert2Html(docPath, dbName, docMdStr)
}

// convert2Html md convert to html
func convert2Html(docPath, dbName, docMdStr string) {
	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_COMPLETE_PAGE
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_USE_XHTML
	renderer := blackfriday.HtmlRenderer(htmlFlags, dbName, "")
	extensions := 0
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_HARD_LINE_BREAK
	extensions |= blackfriday.EXTENSION_HEADER_IDS
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_TABLES

	output := blackfriday.Markdown([]byte(docMdStr), renderer, extensions)
	util.WriteToFile(path.Join(docPath, dbName+".html"), string(output))
	fmt.Println("html generate successfully!")
}
