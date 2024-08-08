package util

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync_k8s_tidb_mysql_data/entity"
)

func CreateDB() *gorm.DB {
	dsn := "root:Mesgxgk@123456@tcp(10.2.65.10:30287)/yifu_produce?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return db
}

// TiDBGormBegin start a TiDB and Gorm transaction as a block. If no error is returned, the transaction will be committed. Otherwise, the transaction will be rolled back.
func TiDBGormBegin(db *gorm.DB, pessimistic bool, fc func(tx *gorm.DB) error) (err error) {
	session := db.Session(&gorm.Session{})
	if session.Error != nil {
		return session.Error
	}

	if pessimistic {
		session = session.Exec("set @@tidb_txn_mode=pessimistic")
	} else {
		session = session.Exec("set @@tidb_txn_mode=optimistic")
	}

	if session.Error != nil {
		return session.Error
	}
	return session.Transaction(fc)
}

func insertBatchRecords(tableName string, records []interface{}) {
	// 初始化数据库连接
	db := CreateDB()

	// 开始事务
	//tx := db.Begin()

	// 每次批量插入的大小
	batchSize := 2000

	// 执行批量插入
	if err := db.Model(&entity.ProduceIn{}).CreateInBatches(records, batchSize).Error; err != nil {
		db.Rollback()
		log.Fatalf("批量插入记录失败: %v", err)
	}

	// 提交事务
	db.Commit()
	fmt.Printf("插入 %d 条记录到数据库表 %s\n", len(records), tableName)
}
