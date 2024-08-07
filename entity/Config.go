package entity

type Tidb struct {
	Url      string
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
