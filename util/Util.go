package util

import (
	"log"
	"strconv"
	"time"
)

// parseInt64 解析 int64
func ParseInt64(s string) int64 {
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
func ParseInt8(s string) int8 {
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
func ParseFloat64(s string) float64 {
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
func ParseTime(s string) time.Time {
	t, err := time.Parse("2006/1/2 15:04:05", s)
	if err != nil {
		log.Printf("解析时间失败: %v", err)
		return time.Time{}
	}
	return t
}

// parseDate 解析日期字符串
func ParseDate(s string) time.Time {
	t, err := time.Parse("2006/1/2", s)
	if err != nil {
		log.Printf("解析日期失败: %v", err)
		return time.Time{}
	}
	return t
}
