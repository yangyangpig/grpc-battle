package sqlx

import "grpc-battle/pkg/dbdriver"

type sql struct {
	op *dbdriver.Opener
}

func NewSqlx() *sql {
	return &sql{op:dbdriver.NewOpener("mysql")}
}

func (x *sql) query(sql string)  {

}

func (x *sql) Insert() {

}

func (x *sql) Update() {

}

func (x *sql) Delete() {

}
