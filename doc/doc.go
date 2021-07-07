package doc

import (
	"db-doc/model"
	"db-doc/util"
	"os"
	"path"
)

// CreateDoc create doc
func CreateDoc(dbInfo model.DbInfo, docType int, tables []model.Table) {
	var docPath string
	dir, _ := os.Getwd()
	if docType == 1 {
		docPath = path.Join(dir, "dist", dbInfo.DbName, "www")
		util.CreateDir(docPath)
		createOnlineDoc(docPath, dbInfo, tables)
	} else {
		docPath = path.Join(dir, "dist", dbInfo.DbName)
		util.CreateDir(docPath)
		createOfflineDoc(docPath, dbInfo, tables)
	}
}
