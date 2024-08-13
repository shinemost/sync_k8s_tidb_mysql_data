package db

import (
	"fmt"
	"log"
	"xorm.io/xorm"
	"xorm.io/xorm/names"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "Mesgxgk@123456"
	HOST     = "10.2.65.10"
	PORT     = 30287
	DATABASE = "yifu_produce"
)

func CreateDB() *xorm.Engine {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&allowAllFiles=true",
		USERNAME,
		PASSWORD,
		HOST,
		PORT,
		DATABASE,
	)

	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 设置数据库表映射规则，使用结构体名称的小写蛇形作为表名
	engine.SetMapper(names.SnakeMapper{})
	engine.ShowSQL(false)

	return engine
}
