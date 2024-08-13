package entity

import (
	"fmt"
	"log"
	"sync_k8s_tidb_mysql_data/db"
	"sync_k8s_tidb_mysql_data/util"
	"time"
)

// ProduceParam 生产工艺参数表
type ProduceParam struct {
	Id                         int64     `json:"id" xorm:"'id' pk autoincr"`                                         // ID
	ProduceId                  int64     `json:"produce_id" xorm:"produce_id"`                                       // 生产ID
	ProductCode                string    `json:"product_code" xorm:"product_code"`                                   // 产品编码
	TechnicsParamId            int64     `json:"technics_param_id" xorm:"technics_param_id"`                         // 工艺参数ID
	TechnicsParamName          string    `json:"technics_param_name" xorm:"technics_param_name"`                     // 工艺参数名称
	TechnicsParamCode          string    `json:"technics_param_code" xorm:"technics_param_code"`                     // 工艺参数编码
	TechnicsParamValue         string    `json:"technics_param_value" xorm:"technics_param_value"`                   // 工艺参数值
	TechnicsParamMaxValue      string    `json:"technics_param_max_value" xorm:"technics_param_max_value"`           // 工艺参数最大值
	TechnicsParamMinValue      string    `json:"technics_param_min_value" xorm:"technics_param_min_value"`           // 工艺参数最小值
	TechnicsParamStandardValue string    `json:"technics_param_standard_value" xorm:"technics_param_standard_value"` // 工艺参数标准/值
	TechnicsParamQuality       int8      `json:"technics_param_quality" xorm:"technics_param_quality"`               // 工艺参数质量：0，不合格；1，合格；
	TechnicsParamType          int64     `json:"technics_param_type" xorm:"technics_param_type"`                     // 工艺参数类型:1:定量,2:定性
	Desc                       string    `json:"desc" xorm:"desc"`                                                   // 描述
	TenantId                   int64     `json:"tenant_id" xorm:"tenant_id"`                                         // 租户ID
	Status                     int8      `json:"status" xorm:"status"`                                               // 状态：0，删除；1，正常；
	CreatedTime                time.Time `json:"created_time" xorm:"created_time"`                                   // 记录添加时间
	LastModifiedTime           time.Time `json:"last_modified_time" xorm:"last_modified_time"`                       // 记录更新时间
}

// TableName 表名称
func (*ProduceParam) TableName() string {
	return "produce_param"
}

func (p *ProduceParam) Parse(line []string) (interface{}, error) {
	return parseProduceParam(line), nil
}

func (p *ProduceParam) InsertBatch(records []interface{}) error {
	var produceRecords []ProduceParam
	for _, r := range records {
		produceRecords = append(produceRecords, r.(ProduceParam))
	}
	return insertProduceParam(produceRecords)
}

// parseProduceParam 解析 ProduceParam 记录
func parseProduceParam(line []string) ProduceParam {
	return ProduceParam{
		Id:                         util.ParseInt64(line[0]),
		ProduceId:                  util.ParseInt64(line[1]),
		ProductCode:                line[2],
		TechnicsParamId:            util.ParseInt64(line[3]),
		TechnicsParamName:          line[4],
		TechnicsParamCode:          line[5],
		TechnicsParamValue:         line[6],
		TechnicsParamMaxValue:      line[7],
		TechnicsParamMinValue:      line[8],
		TechnicsParamStandardValue: line[9],
		TechnicsParamQuality:       util.ParseInt8(line[10]),
		TechnicsParamType:          util.ParseInt64(line[11]),
		Desc:                       line[12],
		TenantId:                   util.ParseInt64(line[13]),
		Status:                     util.ParseInt8(line[14]),
		CreatedTime:                util.ParseTime(line[15]),
		LastModifiedTime:           util.ParseTime(line[16]),
	}
}

func insertProduceParam(records []ProduceParam) error {
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

	fmt.Printf("插入 %d 条记录到数据库表 produce_param 耗时: %v\n", totalRecords, duration)
	return nil
}
