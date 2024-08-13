package entity

import (
	"fmt"
	"log"
	"sync_k8s_tidb_mysql_data/db"
	"sync_k8s_tidb_mysql_data/util"
	"time"
)

// ProduceIn 生产产品投入表
type ProduceIn struct {
	Id               int64     `json:"id" xorm:"'id' pk autoincr"`                   // 自增ID
	ProduceId        int64     `json:"produce_id" xorm:"produce_id"`                 // 生产ID
	ProductCode      string    `json:"product_code" xorm:"product_code"`             // 产品编码
	ProductCount     string    `json:"product_count" xorm:"product_count"`           // 产品数量
	PlanProductCount string    `json:"plan_product_count" xorm:"plan_product_count"` // 计划产品数量
	MaterielCode     string    `json:"materiel_code" xorm:"materiel_code"`           // 物料编码
	MaterielName     string    `json:"materiel_name" xorm:"materiel_name"`           // 物料名称
	MaterielType     int64     `json:"materiel_type" xorm:"materiel_type"`           // 物料类型，1：原材料，2：半成品，3：成品
	ProductOutCode   string    `json:"product_out_code" xorm:"product_out_code"`     // 产品产出编码
	MaterielOutCode  string    `json:"materiel_out_code" xorm:"materiel_out_code"`   // 物料产出编码
	MaterielOutName  string    `json:"materiel_out_name" xorm:"materiel_out_name"`   // 物料产出名称
	MaterielOutType  int64     `json:"materiel_out_type" xorm:"materiel_out_type"`   // 物料产出类型，1：原材料，2：半成品，3：成品
	TenantId         int64     `json:"tenant_id" xorm:"tenant_id"`                   // 租户ID
	Status           int8      `json:"status" xorm:"status"`                         // 状态：0，删除；1，正常；
	CreatedTime      time.Time `json:"created_time" xorm:"created_time"`             // 记录添加时间
	LastModifiedTime time.Time `json:"last_modified_time" xorm:"last_modified_time"` // 记录更新时间
}

// TableName 表名称
func (*ProduceIn) TableName() string {
	return "produce_in"
}

func (p *ProduceIn) Parse(line []string) (interface{}, error) {
	return parseProduceIn(line), nil
}

func (p *ProduceIn) InsertBatch(records []interface{}) error {
	var produceRecords []ProduceIn
	for _, r := range records {
		produceRecords = append(produceRecords, r.(ProduceIn))
	}
	return insertProduceIn(produceRecords)
}

// parseProduceIn 解析 ProduceIn 记录
func parseProduceIn(line []string) ProduceIn {
	return ProduceIn{
		Id:               util.ParseInt64(line[0]),
		ProduceId:        util.ParseInt64(line[1]),
		ProductCode:      line[2],
		ProductCount:     line[3],
		PlanProductCount: line[4],
		MaterielCode:     line[5],
		MaterielName:     line[6],
		MaterielType:     util.ParseInt64(line[7]),
		ProductOutCode:   line[8],
		MaterielOutCode:  line[9],
		MaterielOutName:  line[10],
		MaterielOutType:  util.ParseInt64(line[11]),
		TenantId:         util.ParseInt64(line[12]),
		Status:           util.ParseInt8(line[13]),
		CreatedTime:      util.ParseTime(line[14]),
		LastModifiedTime: util.ParseTime(line[15]),
	}
}

func insertProduceIn(records []ProduceIn) error {
	// 连接数据库
	engine := db.CreateDB()

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
		_, err := engine.Insert(&batch)
		if err != nil {
			log.Fatalf("批量插入失败: %v", err)
			return err
		}
	}

	// 记录结束时间
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Printf("插入 %d 条记录到数据库表 produce_in 耗时: %v\n", totalRecords, duration)
	return nil
}
