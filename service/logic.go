package service

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"sync_k8s_tidb_mysql_data/util"
)

var db *gorm.DB

func Clear(clearTableSlices []string) error {
	db = util.CreateDB()
	tx := db.Begin()

	for _, tableName := range clearTableSlices {
		if err := tx.Exec(fmt.Sprintf("truncate table %s", tableName)).Error; err != nil {
			tx.Rollback()
			log.Fatalf("清除%s表发生错误：%v", tableName, err)
			return err
		}
	}
	tx.Commit()
	return nil
}

// Insert 循环遍历表名，调用 ReadCsv 函数读取数据并处理
func Insert(importTableNames []string) error {
	for _, tableName := range importTableNames {
		if err := util.ReadCsv(tableName); err != nil {
			log.Printf("处理表 %s 时出错: %v", tableName, err)
			return err
		}
	}
	return nil
}

func All() error {
	return nil
}
