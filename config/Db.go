package config

import (
	"database/sql"
	"fmt"
)

func openDB(driverName string, runnable func(db *sql.DB)) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&tls=%s",${tidb_user}, ${tidb_password}, ${tidb_host}, ${tidb_port}, ${tidb_db_name}, ${use_ssl})
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	runnable(db)
}