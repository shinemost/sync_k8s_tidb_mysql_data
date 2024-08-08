package entity

import "time"

// ProuduceParam 生产工艺参数表
type ProuduceParam struct {
	ID                         int64     `json:"id" gorm:"id"`                                                       // ID
	ProduceId                  int64     `json:"produce_id" gorm:"produce_id"`                                       // 生产ID
	ProductCode                string    `json:"product_code" gorm:"product_code"`                                   // 产品编码
	TechnicsParamId            int64     `json:"technics_param_id" gorm:"technics_param_id"`                         // 工艺参数ID
	TechnicsParamName          string    `json:"technics_param_name" gorm:"technics_param_name"`                     // 工艺参数名称
	TechnicsParamCode          string    `json:"technics_param_code" gorm:"technics_param_code"`                     // 工艺参数编码
	TechnicsParamValue         string    `json:"technics_param_value" gorm:"technics_param_value"`                   // 工艺参数值
	TechnicsParamMaxValue      string    `json:"technics_param_max_value" gorm:"technics_param_max_value"`           // 工艺参数最大值
	TechnicsParamMinValue      string    `json:"technics_param_min_value" gorm:"technics_param_min_value"`           // 工艺参数最小值
	TechnicsParamStandardValue string    `json:"technics_param_standard_value" gorm:"technics_param_standard_value"` // 工艺参数标准/值
	TechnicsParamQuality       int8      `json:"technics_param_quality" gorm:"technics_param_quality"`               // 工艺参数质量：0，不合格；1，合格；
	TechnicsParamType          int64     `json:"technics_param_type" gorm:"technics_param_type"`                     // 工艺参数类型:1:定量,2:定性
	Desc                       string    `json:"desc" gorm:"desc"`                                                   // 描述
	TenantId                   int64     `json:"tenant_id" gorm:"tenant_id"`                                         // 租户ID
	Status                     int8      `json:"status" gorm:"status"`                                               // 状态：0，删除；1，正常；
	CreatedTime                time.Time `json:"created_time" gorm:"created_time"`                                   // 记录添加时间
	LastModifiedTime           time.Time `json:"last_modified_time" gorm:"last_modified_time"`                       // 记录更新时间
}

// TableName 表名称
func (*ProuduceParam) TableName() string {
	return "prouduce_param"
}
