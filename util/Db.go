package util

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync_k8s_tidb_mysql_data/cmd"
	"sync_k8s_tidb_mysql_data/entity"
)

func CreateDB() *gorm.DB {
	config := cmd.Config.Tidb

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		CreateBatchSize:        2000,
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
	//Todo 使用tidb的driver是否会有不同
	//ToDo 直接使用load csv的sql语句，看看支不支持
	//https://blog.csdn.net/puss0/article/details/90411618
	// 初始化数据库连接
	db := CreateDB()

	// 开始事务
	//tx := db.Begin()

	// 每次批量插入的大小
	//db.CreateBatchSize = 2000

	// 执行批量插入
	if err := db.Model(&entity.ProduceIn{}).Create(records).Error; err != nil {
		db.Rollback()
		log.Fatalf("批量插入记录失败: %v", err)
	}

	// 提交事务
	db.Commit()
	fmt.Printf("插入 %d 条记录到数据库表 %s\n", len(records), tableName)
}
