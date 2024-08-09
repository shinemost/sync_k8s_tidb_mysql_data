package util

import (
	"fmt"
	"log"
	"sync_k8s_tidb_mysql_data/entity"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"

	_ "github.com/go-sql-driver/mysql"
)

func XCreateDB() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:Mesgxgk@123456@tcp(10.2.65.10:30287)/yifu_producer?charset=utf8")
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 设置数据库表映射规则，使用结构体名称的小写蛇形作为表名
	engine.SetMapper(names.SnakeMapper{})
	engine.ShowSQL(true)

	return engine
}

func insertProduceInX(records []entity.ProduceIn) {
	// 连接数据库
	db := XCreateDB()

	// 记录开始时间
	startTime := time.Now()

	batchSize := 2500
	totalRecords := len(records)

	for i := 0; i < totalRecords; i += batchSize {
		end := i + batchSize
		if end > totalRecords {
			end = totalRecords
		}

		batch := records[i:end]

		// 批量插入
		_, err := db.Insert(&batch)
		if err != nil {
			log.Fatalf("批量插入失败: %v", err)
		}
	}

	// 记录结束时间
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Printf("插入 %d 条记录到数据库表 produce_in 耗时: %v\n", totalRecords, duration)
}
