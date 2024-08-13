package entity

type RecordProcessor interface {
	Parse(line []string) (interface{}, error)
	InsertBatch(records []interface{}) error
}
