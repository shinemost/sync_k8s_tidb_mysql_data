package entity

import (
	"time"
)

// ProduceIn 生产产品投入表
type ProduceIn struct {
	Id               int64     `json:"id" gorm:"id" xorm:"'id' pk autoincr"`                                   // 自增ID
	ProduceId        int64     `json:"produce_id" gorm:"produce_id" xorm:"produce_id"`                         // 生产ID
	ProductCode      string    `json:"product_code" gorm:"product_code"  xorm:"product_code"`                  // 产品编码
	ProductCount     string    `json:"product_count" gorm:"product_count"  xorm:"product_count"`               // 产品数量
	PlanProductCount string    `json:"plan_product_count" gorm:"plan_product_count" xorm:"plan_product_count"` // 计划产品数量
	MaterielCode     string    `json:"materiel_code" gorm:"materiel_code" xorm:"materiel_code"`                // 物料编码
	MaterielName     string    `json:"materiel_name" gorm:"materiel_name" xorm:"materiel_name"`                // 物料名称
	MaterielType     int64     `json:"materiel_type" gorm:"materiel_type" xorm:"materiel_type"`                // 物料类型，1：原材料，2：半成品，3：成品
	ProductOutCode   string    `json:"product_out_code" gorm:"product_out_code" xorm:"product_out_code"`       // 产品产出编码
	MaterielOutCode  string    `json:"materiel_out_code" gorm:"materiel_out_code" xorm:"materiel_out_code"`    // 物料产出编码
	MaterielOutName  string    `json:"materiel_out_name" gorm:"materiel_out_name" xorm:"materiel_out_name"`    // 物料产出名称
	MaterielOutType  int64     `json:"materiel_out_type" gorm:"materiel_out_type" xorm:"materiel_out_type"`    // 物料产出类型，1：原材料，2：半成品，3：成品
	TenantId         int64     `json:"tenant_id" gorm:"tenant_id" xorm:"tenant_id"`                            // 租户ID
	Status           int8      `json:"status" gorm:"status" xorm:"status"`                                     // 状态：0，删除；1，正常；
	CreatedTime      time.Time `json:"created_time" gorm:"created_time" xorm:"created_time"`                   // 记录添加时间
	LastModifiedTime time.Time `json:"last_modified_time" gorm:"last_modified_time" xorm:"last_modified_time"` // 记录更新时间
}

// TableName 表名称
func (*ProduceIn) TableName() string {
	return "produce_in"
}
