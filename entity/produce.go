package entity

import (
	"fmt"
	"log"
	"sync_k8s_tidb_mysql_data/db"
	"sync_k8s_tidb_mysql_data/util"
	"time"
)

// Produce 生产表
type Produce struct {
	Id                  int64     `json:"id" xorm:"'id' pk autoincr"`                         // 自增ID
	Code                string    `json:"code" xorm:"code"`                                   // 生产编码
	ProduceOrderCode    string    `json:"produce_order_code" xorm:"produce_order_code"`       // 生产工单编码
	TechnicsLineId      int64     `json:"technics_line_id" xorm:"technics_line_id"`           // 工艺ID
	TechnicsLineCode    string    `json:"technics_line_code" xorm:"technics_line_code"`       // 工艺编码
	TechnicsLineName    string    `json:"technics_line_name" xorm:"technics_line_name"`       // 工艺名称
	TechnicsProcessId   int64     `json:"technics_process_id" xorm:"technics_process_id"`     // 工序ID
	TechnicsProcessCode string    `json:"technics_process_code" xorm:"technics_process_code"` // 工序编码
	TechnicsProcessName string    `json:"technics_process_name" xorm:"technics_process_name"` // 工序名称
	TechnicsStepCode    string    `json:"technics_step_code" xorm:"technics_step_code"`       // 工步编码
	TechnicsStepName    string    `json:"technics_step_name" xorm:"technics_step_name"`       // 工步名称
	ProductCode         string    `json:"product_code" xorm:"product_code"`                   // 产品编码
	ProductCount        string    `json:"product_count" xorm:"product_count"`                 // 产品数量
	OkCount             string    `json:"ok_count" xorm:"ok_count"`                           // 合格数量
	NgCount             string    `json:"ng_count" xorm:"ng_count"`                           // 不合格数量
	ProductQuality      int8      `json:"product_quality" xorm:"product_quality"`             // 产品质量：0，不合格；1，合格；
	QualityNgCount      float64   `json:"quality_ng_count" xorm:"quality_ng_count"`
	QualityQuality      int8      `json:"quality_quality" xorm:"quality_quality"`       // 1,合格；2，不合格；
	MaterielCode        string    `json:"materiel_code" xorm:"materiel_code"`           // 物料编码
	MaterielName        string    `json:"materiel_name" xorm:"materiel_name"`           // 物料名称
	MaterielType        int64     `json:"materiel_type"  xorm:"materiel_type"`          // 物料类型
	ProduceDate         time.Time `json:"produce_date" xorm:"produce_date"`             // 生产日期
	StartTime           time.Time `json:"start_time" xorm:"start_time"`                 // 生产开始时间
	EndTime             time.Time `json:"end_time" xorm:"end_time"`                     // 生产结束时间
	UserId              int64     `json:"user_id" xorm:"user_id"`                       // 用户ID
	UserName            string    `json:"user_name" xorm:"user_name"`                   // 用户名称
	UserAccount         string    `json:"user_account" xorm:"user_account"`             // 用户账号
	DeviceId            int64     `json:"device_id" xorm:"device_id"`                   // 设备ID
	DeviceCode          string    `json:"device_code" xorm:"device_code"`               // 设备编码
	DeviceName          string    `json:"device_name" xorm:"device_name"`               // 设备名称
	Remarks             string    `json:"remarks" xorm:"remarks"`                       // 备注
	TenantId            int64     `json:"tenant_id" xorm:"tenant_id"`                   // 租户ID
	Status              int8      `json:"status" xorm:"status"`                         // 状态：0，删除；1，正常；2，结束；-1，解绑；-2，报废；-3，降级；-4，让步放行；-5，返工;
	CreatedTime         time.Time `json:"created_time" xorm:"created_time"`             // 记录添加时间
	LastModifiedTime    time.Time `json:"last_modified_time" xorm:"last_modified_time"` // 记录更新时间
}

// TableName 表名称
func (*Produce) TableName() string {
	return "produce"
}

func (p *Produce) Parse(line []string) (interface{}, error) {
	return parseProduce(line), nil
}

func (p *Produce) InsertBatch(records []interface{}) error {
	var produceRecords []Produce
	for _, r := range records {
		produceRecords = append(produceRecords, r.(Produce))
	}
	return insertProduce(produceRecords)
}

// parseProduce 解析 Produce 记录
func parseProduce(line []string) Produce {
	return Produce{
		Id:                  util.ParseInt64(line[0]),
		Code:                line[1],
		ProduceOrderCode:    line[2],
		TechnicsLineId:      util.ParseInt64(line[3]),
		TechnicsLineCode:    line[4],
		TechnicsLineName:    line[5],
		TechnicsProcessId:   util.ParseInt64(line[6]),
		TechnicsProcessCode: line[7],
		TechnicsProcessName: line[8],
		TechnicsStepCode:    line[9],
		TechnicsStepName:    line[10],
		ProductCode:         line[11],
		ProductCount:        line[12],
		OkCount:             line[13],
		NgCount:             line[14],
		ProductQuality:      util.ParseInt8(line[15]),
		QualityNgCount:      util.ParseFloat64(line[16]),
		QualityQuality:      util.ParseInt8(line[17]),
		MaterielCode:        line[18],
		MaterielName:        line[19],
		MaterielType:        util.ParseInt64(line[20]),
		ProduceDate:         util.ParseDate(line[21]),
		StartTime:           util.ParseTime(line[22]),
		EndTime:             util.ParseTime(line[23]),
		UserId:              util.ParseInt64(line[24]),
		UserName:            line[25],
		UserAccount:         line[26],
		DeviceId:            util.ParseInt64(line[27]),
		DeviceCode:          line[28],
		DeviceName:          line[29],
		Remarks:             line[30],
		TenantId:            util.ParseInt64(line[31]),
		Status:              util.ParseInt8(line[32]),
		CreatedTime:         util.ParseTime(line[33]),
		LastModifiedTime:    util.ParseTime(line[34]),
	}
}

func insertProduce(records []Produce) error {
	// 连接数据库
	engine := db.CreateDB()

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
		_, err := engine.Insert(&batch)
		if err != nil {
			log.Fatalf("批量插入失败: %v", err)
			return err
		}
	}

	// 记录结束时间
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Printf("插入 %d 条记录到数据库表 produce 耗时: %v\n", totalRecords, duration)
	return nil
}
