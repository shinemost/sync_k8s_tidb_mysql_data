package service

import (
	"fmt"
	"gorm.io/gorm"
	"sync_k8s_tidb_mysql_data/util"
)

var db *gorm.DB

func Clear(clearTableSlices []string) error {
	db = util.CreateDB()
	defer db.Commit()

	for _, tableName := range clearTableSlices {
		db.Exec(fmt.Sprintf("truncate table %s", tableName))
	}
	return nil
}

func Insert() error {
	return nil
}

func All() error {
	return nil
}
