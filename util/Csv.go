package util

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync_k8s_tidb_mysql_data/entity"
	"time"
)

func ReadCsv(tableName string) error {
	// 打开 CSV 文件
	file, err := os.Open(fmt.Sprintf("data/%s_0801.csv", tableName))
	if err != nil {
		return fmt.Errorf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 创建一个 CSV Reader
	reader := csv.NewReader(file)

	// 读取并丢弃表头
	if _, err := reader.Read(); err != nil {
		if err.Error() == "EOF" {
			return fmt.Errorf("CSV 文件为空，没有数据")
		}
		return fmt.Errorf("读取 CSV 表头失败: %v", err)
	}

	// 根据表名选择合适的切片
	switch tableName {
	case "produce":
		var records []entity.Produce
		for {
			line, err := reader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					break // 结束循环，已读取完所有数据
				}
				return fmt.Errorf("读取 CSV 行失败: %v", err)
			}
			records = append(records, parseProduce(line))
		}
		if err := insertBatchRecords(tableName, records, len(records)); err != nil {
			return fmt.Errorf("批量插入 Produce 记录失败: %v", err)
		}
	case "produce_param":
		var records []entity.ProduceParam
		for {
			line, err := reader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					break // 结束循环，已读取完所有数据
				}
				return fmt.Errorf("读取 CSV 行失败: %v", err)
			}
			records = append(records, parseProduceParam(line))
		}

		if err := insertBatchRecords(tableName, records, len(records)); err != nil {
			return fmt.Errorf("批量插入 ProduceParam 记录失败: %v", err)
		}
	case "produce_in":
		var records []entity.ProduceIn
		for {
			line, err := reader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					break // 结束循环，已读取完所有数据
				}
				return fmt.Errorf("读取 CSV 行失败: %v", err)
			}
			records = append(records, parseProduceIn(line))
		}
		if err := insertBatchRecords(tableName, records, len(records)); err != nil {
			return fmt.Errorf("批量插入 ProduceIn 记录失败: %v", err)
		}
	default:
		return fmt.Errorf("未知的表名: %s", tableName)
	}
	return nil
}

// parseProduce 解析 Produce 记录
func parseProduce(line []string) entity.Produce {
	return entity.Produce{
		Id:                  parseInt64(line[0]),
		Code:                line[1],
		ProduceOrderCode:    line[2],
		TechnicsLineId:      parseInt64(line[3]),
		TechnicsLineCode:    line[4],
		TechnicsLineName:    line[5],
		TechnicsProcessId:   parseInt64(line[6]),
		TechnicsProcessCode: line[7],
		TechnicsProcessName: line[8],
		TechnicsStepCode:    line[9],
		TechnicsStepName:    line[10],
		ProductCode:         line[11],
		ProductCount:        line[12],
		OkCount:             line[13],
		NgCount:             line[14],
		ProductQuality:      parseInt8(line[15]),
		QualityNgCount:      parseFloat64(line[16]),
		QualityQuality:      parseInt8(line[17]),
		MaterielCode:        line[18],
		MaterielName:        line[19],
		MaterielType:        parseInt64(line[20]),
		ProduceDate:         parseTime(line[21]),
		StartTime:           parseTime(line[22]),
		EndTime:             parseTime(line[23]),
		UserId:              parseInt64(line[24]),
		UserName:            line[25],
		UserAccount:         line[26],
		DeviceId:            parseInt64(line[27]),
		DeviceCode:          line[28],
		DeviceName:          line[29],
		Remarks:             line[30],
		TenantId:            parseInt64(line[31]),
		Status:              parseInt8(line[32]),
		CreatedTime:         parseTime(line[33]),
		LastModifiedTime:    parseTime(line[34]),
	}
}

// parseProduceParam 解析 ProduceParam 记录
func parseProduceParam(line []string) entity.ProduceParam {
	return entity.ProduceParam{
		Id:                         parseInt64(line[0]),
		ProduceId:                  parseInt64(line[1]),
		ProductCode:                line[2],
		TechnicsParamId:            parseInt64(line[3]),
		TechnicsParamName:          line[4],
		TechnicsParamCode:          line[5],
		TechnicsParamValue:         line[6],
		TechnicsParamMaxValue:      line[7],
		TechnicsParamMinValue:      line[8],
		TechnicsParamStandardValue: line[9],
		TechnicsParamQuality:       parseInt8(line[10]),
		TechnicsParamType:          parseInt64(line[11]),
		Desc:                       line[12],
		TenantId:                   parseInt64(line[13]),
		Status:                     parseInt8(line[14]),
		CreatedTime:                parseTime(line[15]),
		LastModifiedTime:           parseTime(line[16]),
	}
}

// parseProduceIn 解析 ProduceIn 记录
func parseProduceIn(line []string) entity.ProduceIn {
	return entity.ProduceIn{
		Id:               parseInt64(line[0]),
		ProduceId:        parseInt64(line[1]),
		ProductCode:      line[2],
		ProductCount:     line[3],
		PlanProductCount: line[4],
		MaterielCode:     line[5],
		MaterielName:     line[6],
		MaterielType:     parseInt64(line[7]),
		ProductOutCode:   line[8],
		MaterielOutCode:  line[9],
		MaterielOutName:  line[10],
		MaterielOutType:  parseInt64(line[11]),
		TenantId:         parseInt64(line[12]),
		Status:           parseInt8(line[13]),
		CreatedTime:      parseTime(line[14]),
		LastModifiedTime: parseTime(line[15]),
	}
}

// parseInt64 解析 int64
func parseInt64(s string) int64 {
	if s == "" {
		return 0 // 或者选择其他默认值
	}
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Printf("解析 int64 失败: %v", err)
		return 0
	}
	return val
}

// parseInt8 解析 int8
func parseInt8(s string) int8 {
	if s == "" {
		return 0 // 或者选择其他默认值
	}
	val, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		log.Printf("解析 int8 失败: %v", err)
		return 0
	}
	return int8(val)
}

// parseFloat64 解析 float64
func parseFloat64(s string) float64 {
	if s == "" {
		return 0 // 或者选择其他默认值
	}
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("解析 float64 失败: %v", err)
		return 0
	}
	return val
}

// parseTime 解析时间字符串
func parseTime(s string) time.Time {
	t, err := time.Parse("2006/1/2 15:04:05", s)
	if err != nil {
		log.Printf("解析时间失败: %v", err)
		return time.Time{}
	}
	return t
}
