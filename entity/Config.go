package entity

type Tidb struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

type Minio struct {
	Url     string
	Bucket  string
	CsvName string
}

type Config struct {
	Tidb  Tidb
	Minio Minio
}
