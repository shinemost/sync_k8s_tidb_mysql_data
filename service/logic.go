package service

import (
	"fmt"
	"gorm.io/gorm"
	"sync_k8s_tidb_mysql_data/util"
)

var db *gorm.DB

func Clear(clearTableSlices []string) error {
	db = util.CreateDB()
	tx := db.Begin()
	defer tx.Commit()

	for _, tableName := range clearTableSlices {
		tx.Exec(fmt.Sprintf("truncate table %s", tableName))
	}
	return nil
}

func Insert(importTableNames []string) error {
	util.ReadCsv(importTableNames[0])
	return nil
}

func All() error {
	return nil
}
