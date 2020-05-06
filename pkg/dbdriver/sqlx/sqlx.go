package sqlx

import "grpc-battle/pkg/dbdriver"

type Sqlx struct {
	op *dbdriver.Opener
}

func NewSqlx() *Sqlx {
	return &Sqlx{op:dbdriver.NewOpener("mysql")}
}

func (x Sqlx) query()  {

}

func (x Sqlx) Insert() {

}

func (x Sqlx) Update() {

}

func (x Sqlx) Delete() {

}
