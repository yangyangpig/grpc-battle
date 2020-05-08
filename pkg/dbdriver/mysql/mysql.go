package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"grpc-battle/pkg/dbdriver"
	"log"
)

var (
	ERR_NOTMYSQL = errors.New("is not mysqlx driver please check again")
)


type mysql struct {
	op *sqlx.DB
}

func NewSqlx() (*mysql, error) {
	dr, err := dbdriver.NewOpener("mysql", dbdriver.MYSQL)
	if err != nil {
		log.Fatalf("newSqlx mysql fail %+v", err)
		return nil, err
	}
	if _, ok := dr.Driver.(*sqlx.DB); !ok {
		return nil, ERR_NOTMYSQL
	}
	mysqlDriver := dr.Driver.(*sqlx.DB)

	return &mysql{op: mysqlDriver}, nil
}

// query one row
func (x *mysql) QueryRow(query string, param ...interface{}) (*Scanner, error) {
	query, args, err := sqlx.In(query, param)
	if err != nil {
		return nil, err
	}

	rows := x.op.QueryRowxContext(context.Background(), query, args...)
	return &Scanner{rowx: rows}, nil
}

// query more than one row
func (x *mysql) QueryRows(query string, param ...interface{}) (*Scanner, error) {
	// bind param
	q, args, err := sqlx.In(query, param)
	if err != nil {
		return nil, err
	}

	rows, err := x.op.QueryxContext(context.Background(), q, args...)
	if err != nil {
		return nil, err
	}
	return &Scanner{rowx: rows}, nil
}

func (x *mysql) ExecSql(query string, param ...interface{}) (*Scanner, error) {

	result, err := x.op.ExecContext(context.Background(), query, param...)

	if err != nil {
		return nil, err
	}
	return &Scanner{rowx: result}, nil
}

// transaction
func (x *mysql) Transaction(handle func(interface{}) error) error {
	tx, err := x.op.Beginx()
	if nil != err {
		return err
	}
	err = handle(tx)

	if nil != err {
		_ = tx.Rollback()
	} else {
		err = tx.Commit()
	}
	return err
}

type Scanner struct {
	rowx interface{}
}

func (s *Scanner) scanner() (res interface{}, err error) {
	switch r := s.rowx.(type) {
	case *sqlx.Rows:
		var m = map[string]interface{}{}
		for r.Next() {
			if err = r.MapScan(m); err == nil {
				fmt.Printf("scan value %+v \n", m)
			}
		}
		// 不关闭连接，会造成连接池膨胀
		defer r.Close()
		res = m
		return

	case *sqlx.Row:
	// TODO

	case sql.Result:
		rowsAffected, _ := r.RowsAffected()

		return rowsAffected, nil
	}
	return
}

// TODO 批处理
func (x *mysql) BatchExec(query string, param ...interface{}) (*Scanner, error) {
	stmt, err := x.op.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	scan, err := stmt.ExecContext(context.Background(), param)

	if err != nil {
		return nil, err
	}

	return &Scanner{rowx:scan}, nil

}