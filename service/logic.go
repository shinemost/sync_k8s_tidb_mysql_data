package service

import (
	"fmt"
	"log"
	"sync"
	"sync_k8s_tidb_mysql_data/db"
)

func Clear(clearTableSlices []string) error {
	engine := db.CreateDB()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}

	for _, tableName := range clearTableSlices {
		if _, err := session.Exec(fmt.Sprintf("truncate table %s", tableName)); err != nil {
			_ = session.Rollback()
			log.Fatalf("清除%s表发生错误：%v", tableName, err)
			return err
		}
	}
	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}

// Insert 循环遍历表名，调用 ReadCsv 函数读取数据并处理
func Insert(importTableNames []string) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(importTableNames))

	for _, tableName := range importTableNames {
		wg.Add(1)
		go func(tName string) {
			defer wg.Done()
			if err := readCsv(tName); err != nil {
				errCh <- fmt.Errorf("处理表 %s 时出错: %v", tName, err)
			}
		}(tableName)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		log.Println(err)
		return err
	}
	return nil
}

// All 先清理再导入
func All(tableNames []string) error {
	err := Clear(tableNames)
	if err != nil {
		return err
	}
	err = Insert(tableNames)
	if err != nil {
		return err
	}
	return nil
}
