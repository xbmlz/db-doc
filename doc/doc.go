package doc

import (
	"db-doc/model"
	"db-doc/util"
	"os"
	"path"
)

// CreateDoc create doc
func CreateDoc(dbName string, docType int, tables []model.Table) {
	var docPath string
	dir, _ := os.Getwd()
	if docType == 1 {
		docPath = path.Join(dir, dbName, "online")
		util.CreateDir(docPath)
		createOnlineDoc(docPath, dbName, tables)
	} else {
		docPath = path.Join(dir, dbName, "offline")
		util.CreateDir(docPath)
		createOfflineDoc()
	}
}
