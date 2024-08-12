package util

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"sync_k8s_tidb_mysql_data/entity"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDB() *xorm.Engine {
	config := entity.Config{}

	// 读取配置
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&allowAllFiles=true",
		config.Tidb.Username,
		config.Tidb.Password,
		config.Tidb.Host,
		config.Tidb.Port,
		config.Tidb.Database,
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

// insertBatchRecords 根据表名调用不同的批量插入方法
func insertBatchRecords(tableName string, records interface{}, nums int) error {
	switch tableName {
	case "produce":
		if recs, ok := records.([]entity.Produce); ok {
			insertProduceX(recs)
		} else {
			return fmt.Errorf("记录类型不匹配: expected []entity.Produce, got %T", records)
		}
	case "produce_param":
		if recs, ok := records.([]entity.ProduceParam); ok {
			insertProduceParamX(recs)
		} else {
			return fmt.Errorf("记录类型不匹配: expected []entity.ProduceParam, got %T", records)
		}
	case "produce_in":
		if recs, ok := records.([]entity.ProduceIn); ok {
			insertProduceInX(recs)
		} else {
			return fmt.Errorf("记录类型不匹配: expected []entity.ProduceIn, got %T", records)
		}
	default:
		return fmt.Errorf("未知的表名: %s", tableName)
	}
	return nil
}

func insertProduceInX(records []entity.ProduceIn) {
	// 连接数据库
	db := CreateDB()

	// 记录开始时间
	startTime := time.Now()

	batchSize := 3000
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

func insertProduceX(records []entity.Produce) {
	// 连接数据库
	db := CreateDB()

	// 记录开始时间
	startTime := time.Now()

	batchSize := 1500
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

	fmt.Printf("插入 %d 条记录到数据库表 produce 耗时: %v\n", totalRecords, duration)
}

func insertProduceParamX(records []entity.ProduceParam) {
	// 连接数据库
	db := CreateDB()

	// 记录开始时间
	startTime := time.Now()

	batchSize := 3000
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

	fmt.Printf("插入 %d 条记录到数据库表 produce_param 耗时: %v\n", totalRecords, duration)
}
