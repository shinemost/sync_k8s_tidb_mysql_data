package util

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync_k8s_tidb_mysql_data/entity"
	"time"
)

func ReadCsv(tableName string) {
	// 打开 CSV 文件
	file, err := os.Open(fmt.Sprintf("data/%s_0801.csv", tableName))
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 创建一个 CSV Reader
	reader := csv.NewReader(file)

	// 读取第一行，获取字段名
	header, err := reader.Read()
	if err != nil {
		log.Fatalf("读取 CSV 头部失败: %v", err)
	}

	// 根据表名选择对应的结构体类型
	var recordType reflect.Type
	switch tableName {
	case "produce":
		recordType = reflect.TypeOf(entity.Produce{})
	case "produce_in":
		recordType = reflect.TypeOf(entity.ProduceIn{})
	case "produce_param":
		recordType = reflect.TypeOf(entity.ProduceParam{})
	default:
		log.Fatalf("未知的表名: %s", tableName)
	}

	// 解析数据并映射到结构体
	var records []interface{}
	for {
		line, err := reader.Read()
		if err != nil {
			break // 结束循环，已读取完所有数据或发生错误
		}

		record := reflect.New(recordType).Interface()
		// 使用反射设置结构体字段值
		err = setField(record, header, line)
		if err != nil {
			log.Fatalf("设置字段值失败: %v", err)
		}

		records = append(records, record)
	}

	// 打印解析的记录（可选）
	fmt.Printf("读取到 %d 条记录\n", len(records))

	// 批量插入到数据库

	// 调用批量插入函数
	insertBatchRecords(tableName, records)

}

// setField 使用反射设置结构体字段值
func setField(record interface{}, header []string, line []string) error {
	recordValue := reflect.ValueOf(record).Elem()
	// 遍历 CSV 的每一列
	for i, fieldName := range header {
		// 转换字段名格式：从下划线分隔转为驼峰式
		structFieldName := toCamelCase(fieldName)
		// 获取结构体字段
		field := recordValue.FieldByName(structFieldName)
		if !field.IsValid() {
			return fmt.Errorf("未找到结构体字段: %s", fieldName)
		}

		// 解析并设置字段值
		// 解析并设置字段值
		switch field.Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if line[i] == "" {
				field.SetInt(0)
			} else {
				val, err := strconv.ParseInt(line[i], 10, 64)
				if err != nil {
					return fmt.Errorf("转换为整数失败: %v", err)
				}
				field.SetInt(val)
			}
		case reflect.String:
			field.SetString(line[i])
		case reflect.Struct:
			if field.Type() == reflect.TypeOf(time.Time{}) {
				if line[i] == "" {
					field.Set(reflect.ValueOf(time.Now()))
				} else {
					// 解析时间字符串
					t, err := parseTime(line[i])
					if err != nil {
						return fmt.Errorf("解析时间失败: %v", err)
					}
					field.Set(reflect.ValueOf(t))
				}
			} else {
				return fmt.Errorf("不支持的结构体类型: %v", field.Type())
			}
		default:
			return fmt.Errorf("不支持的字段类型: %v", field.Type().Kind())
		}
	}

	return nil
}

// toCamelCase 将下划线分隔的字符串转换为驼峰式
func toCamelCase(s string) string {
	// 将下划线分隔的字符串转为驼峰式
	parts := strings.Split(s, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, "")
}

// parseTime 解析时间字符串
func parseTime(s string) (time.Time, error) {
	return time.Parse("2006/1/2 15:04:05", s)
}
