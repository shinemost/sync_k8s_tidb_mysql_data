package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync_k8s_tidb_mysql_data/entity"
)

func getProcessor(tableName string) entity.RecordProcessor {
	switch tableName {
	case "produce":
		return &entity.Produce{}
	case "produce_param":
		return &entity.ProduceParam{}
	case "produce_in":
		return &entity.ProduceIn{}
	default:
		return nil
	}
}

func readCsv(tableName string) error {
	processor := getProcessor(tableName)
	if processor == nil {
		return fmt.Errorf("未知的表名: %s", tableName)
	}

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

	var records []interface{}
	for {
		line, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break // 结束循环，已读取完所有数据
			}
			return fmt.Errorf("读取 CSV 行失败: %v", err)
		}
		record, err := processor.Parse(line)
		if err != nil {
			return fmt.Errorf("解析记录失败: %v", err)
		}
		records = append(records, record)
	}

	if err := processor.InsertBatch(records); err != nil {
		return fmt.Errorf("批量插入记录失败: %v", err)
	}

	return nil
}
