package core

type Table struct {
	TableName string `json:"table_name"`
	Rows      []*Row `json:"rows"`
}
