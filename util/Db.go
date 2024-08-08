package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
