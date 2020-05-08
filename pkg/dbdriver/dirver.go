package dbdriver

type Driver interface {
	ExecSql(query string, param ...interface{}) (interface{}, error)
	QueryRow(query string, param ...interface{}) (interface{}, error)
	QueryRows(query string, param ...interface{}) (interface{}, error)
	Transaction(handle func(interface{}) error) error
}
