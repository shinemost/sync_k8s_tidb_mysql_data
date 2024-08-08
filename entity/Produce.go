package entity

import "time"

// Produce 生产表
type Produce struct {
	Id                  int64     `json:"id" gorm:"id"`                                       // 自增ID
	Code                string    `json:"code" gorm:"code"`                                   // 生产编码
	ProduceOrderCode    string    `json:"produce_order_code" gorm:"produce_order_code"`       // 生产工单编码
	TechnicsLineId      int64     `json:"technics_line_id" gorm:"technics_line_id"`           // 工艺ID
	TechnicsLineCode    string    `json:"technics_line_code" gorm:"technics_line_code"`       // 工艺编码
	TechnicsLineName    string    `json:"technics_line_name" gorm:"technics_line_name"`       // 工艺名称
	TechnicsProcessId   int64     `json:"technics_process_id" gorm:"technics_process_id"`     // 工序ID
	TechnicsProcessCode string    `json:"technics_process_code" gorm:"technics_process_code"` // 工序编码
	TechnicsProcessName string    `json:"technics_process_name" gorm:"technics_process_name"` // 工序名称
	TechnicsStepCode    string    `json:"technics_step_code" gorm:"technics_step_code"`       // 工步编码
	TechnicsStepName    string    `json:"technics_step_name" gorm:"technics_step_name"`       // 工步名称
	ProductCode         string    `json:"product_code" gorm:"product_code"`                   // 产品编码
	ProductCount        string    `json:"product_count" gorm:"product_count"`                 // 产品数量
	OkCount             string    `json:"ok_count" gorm:"ok_count"`                           // 合格数量
	NgCount             string    `json:"ng_count" gorm:"ng_count"`                           // 不合格数量
	ProductQuality      int8      `json:"product_quality" gorm:"product_quality"`             // 产品质量：0，不合格；1，合格；
	QualityNgCount      float64   `json:"quality_ng_count" gorm:"quality_ng_count"`
	QualityQuality      int8      `json:"quality_quality" gorm:"quality_quality"`       // 1,合格；2，不合格；
	MaterielCode        string    `json:"materiel_code" gorm:"materiel_code"`           // 物料编码
	MaterielName        string    `json:"materiel_name" gorm:"materiel_name"`           // 物料名称
	MaterielType        int64     `json:"materiel_type" gorm:"materiel_type"`           // 物料类型
	ProduceDate         time.Time `json:"produce_date" gorm:"produce_date"`             // 生产日期
	StartTime           time.Time `json:"start_time" gorm:"start_time"`                 // 生产开始时间
	EndTime             time.Time `json:"end_time" gorm:"end_time"`                     // 生产结束时间
	UserId              int64     `json:"user_id" gorm:"user_id"`                       // 用户ID
	UserName            string    `json:"user_name" gorm:"user_name"`                   // 用户名称
	UserAccount         string    `json:"user_account" gorm:"user_account"`             // 用户账号
	DeviceId            int64     `json:"device_id" gorm:"device_id"`                   // 设备ID
	DeviceCode          string    `json:"device_code" gorm:"device_code"`               // 设备编码
	DeviceName          string    `json:"device_name" gorm:"device_name"`               // 设备名称
	Remarks             string    `json:"remarks" gorm:"remarks"`                       // 备注
	TenantId            int64     `json:"tenant_id" gorm:"tenant_id"`                   // 租户ID
	Status              int8      `json:"status" gorm:"status"`                         // 状态：0，删除；1，正常；2，结束；-1，解绑；-2，报废；-3，降级；-4，让步放行；-5，返工;
	CreatedTime         time.Time `json:"created_time" gorm:"created_time"`             // 记录添加时间
	LastModifiedTime    time.Time `json:"last_modified_time" gorm:"last_modified_time"` // 记录更新时间
}

// TableName 表名称
func (*Produce) TableName() string {
	return "produce"
}
