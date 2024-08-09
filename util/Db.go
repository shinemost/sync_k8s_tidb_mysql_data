package util

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"sync_k8s_tidb_mysql_data/entity"
	"time"
)

func CreateDB() *gorm.DB {
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
		//SkipDefaultTransaction: true,
		CreateBatchSize: 2000,
		PrepareStmt:     true,
	})
	if err != nil {
		panic(err)
	}

	sqlDb, _ := db.DB()

	sqlDb.SetMaxIdleConns(100)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(5 * time.Minute)

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

// insertProduce 批量插入 Produce 记录，每2000条提交一次事务
func insertProduce(records []entity.Produce) {
	db := CreateDB()
	//defer db.Close()

	// 记录开始时间
	startTime := time.Now()

	batchSize := 2000
	totalRecords := len(records)
	var wg sync.WaitGroup
	sem := make(chan struct{}, 14) // 限制并发数为10

	processBatch := func(batch []entity.Produce) {
		defer wg.Done()
		defer func() { <-sem }() // Release semaphore slot

		// 创建事务
		tx := db.Begin()
		if tx.Error != nil {
			log.Printf("开始事务失败: %v", tx.Error)
			return
		}

		// 执行批量插入
		err := tx.Model(&entity.Produce{}).Create(batch).Error
		if err != nil {
			tx.Rollback()
			log.Printf("批量插入记录失败: %v", err)
			return
		}

		// 提交事务
		if err := tx.Commit().Error; err != nil {
			log.Printf("提交事务失败: %v", err)
		}
	}

	for i := 0; i < totalRecords; i += batchSize {
		end := i + batchSize
		if end > totalRecords {
			end = totalRecords
		}

		batch := records[i:end]
		wg.Add(1)
		sem <- struct{}{} // Acquire a semaphore slot
		go processBatch(batch)
	}

	wg.Wait()

	// 记录结束时间
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Printf("插入 %d 条记录到数据库表 produce 耗时: %v\n", totalRecords, duration)
}

// insertProduceParam 批量插入 ProduceParam 记录，每2000条提交一次事务
func insertProduceParam(records []entity.ProduceParam) {
	db := CreateDB()
	//defer db.Close()

	// 记录开始时间
	startTime := time.Now()

	batchSize := 2000
	totalRecords := len(records)
	var wg sync.WaitGroup
	sem := make(chan struct{}, 10) // 限制并发数为10

	processBatch := func(batch []entity.ProduceParam) {
		defer wg.Done()
		defer func() { <-sem }() // Release semaphore slot

		// 创建事务
		tx := db.Begin()
		if tx.Error != nil {
			log.Printf("开始事务失败: %v", tx.Error)
			return
		}

		// 执行批量插入
		err := tx.Model(&entity.ProduceParam{}).Create(batch).Error
		if err != nil {
			tx.Rollback()
			log.Printf("批量插入记录失败: %v", err)
			return
		}

		// 提交事务
		if err := tx.Commit().Error; err != nil {
			log.Printf("提交事务失败: %v", err)
		}
	}

	for i := 0; i < totalRecords; i += batchSize {
		end := i + batchSize
		if end > totalRecords {
			end = totalRecords
		}

		batch := records[i:end]
		wg.Add(1)
		sem <- struct{}{} // Acquire a semaphore slot
		go processBatch(batch)
	}

	wg.Wait()

	// 记录结束时间
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Printf("插入 %d 条记录到数据库表 produce_param 耗时: %v\n", totalRecords, duration)
}

// insertProduceIn 批量插入 ProduceIn 记录，每2000条提交一次事务
func insertProduceIn(records []entity.ProduceIn) {
	// 连接数据库
	db := CreateDB()

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

		// 构造 SQL 插入语句
		query := "INSERT INTO produce_in (id, produce_id, product_code, product_count, plan_product_count, materiel_code, materiel_name, materiel_type, product_out_code, materiel_out_code, materiel_out_name, materiel_out_type, tenant_id, status, created_time, last_modified_time) VALUES "
		values := []interface{}{}

		for j, record := range batch {
			if j > 0 {
				query += ", "
			}
			query += "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
			values = append(values,
				record.Id, record.ProduceId, record.ProductCode, record.ProductCount, record.PlanProductCount,
				record.MaterielCode, record.MaterielName, record.MaterielType, record.ProductOutCode,
				record.MaterielOutCode, record.MaterielOutName, record.MaterielOutType,
				record.TenantId, record.Status, record.CreatedTime, record.LastModifiedTime,
			)
		}

		err := db.Exec(query, values...).Error
		if err != nil {
			log.Fatalf("批量插入失败: %v", err)
		}

	}

	// 记录结束时间
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Printf("插入 %d 条记录到数据库表 produce_in 耗时: %v\n", totalRecords, duration)
}

// insertBatchRecords 根据表名调用不同的批量插入方法
func insertBatchRecords(tableName string, records interface{}, nums int) error {
	switch tableName {
	case "produce":
		if recs, ok := records.([]entity.Produce); ok {
			insertProduce(recs)
		} else {
			return fmt.Errorf("记录类型不匹配: expected []entity.Produce, got %T", records)
		}
	case "produce_param":
		if recs, ok := records.([]entity.ProduceParam); ok {
			insertProduceParam(recs)
		} else {
			return fmt.Errorf("记录类型不匹配: expected []entity.ProduceParam, got %T", records)
		}
	case "produce_in":
		if recs, ok := records.([]entity.ProduceIn); ok {
			insertProduceIn(recs)
		} else {
			return fmt.Errorf("记录类型不匹配: expected []entity.ProduceIn, got %T", records)
		}
	default:
		return fmt.Errorf("未知的表名: %s", tableName)
	}
	return nil
}
