package dbdriver

type DBDriver interface {
	GetDB() (data []Db, err error)
	GetTables(dbName string) (data []Table, err error)
	GetTableColumns(dbName string, tableName string) (data []Column, err error)
}
