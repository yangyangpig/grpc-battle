package mysql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)

const (
	userRemark   = "to_uid, create_at, dropped, id, sequence_id"
	insertRemark = "to_uid, create_at"
)

var sqlDriver *mysql

func init() {
	sqlDriver, _ = NewSqlx()
}
func TestSql_QueryRows(t *testing.T) {
	// 如果不设置最大打开连接数,不限制，就不断创建新连接
	// sqlDriver.op.SetMaxOpenConns(1)
	query := fmt.Sprintf("select %s from user_friendships where from_uid in (?)", userRemark)
	scan, err := sqlDriver.QueryRows(query, 500001, 500001)
	if err != nil {
		t.Errorf("queryrows fail %v", err)
		return
	}

	// 模拟阻塞的情况，由于设置了最大的打开连接数为1，而上面的query没有关闭conn，所以下面就会一直阻塞,如果不设置最大打开连接数,不限制，就不断创建新连接
	query2 := fmt.Sprintf("select %s from user_friendships where from_uid in (?)", userRemark)
	_, err = sqlDriver.QueryRows(query2, 500001, 500001)

	if err != nil {
		t.Errorf("queryrows2 fail %v", err)
		return
	}

	m, err := scan.scanner()
	if err != nil {
		t.Errorf("scanner fail %v", err)
		return
	}

	fmt.Printf("last result %+v", m)

}

func TestMysql_Insert(t *testing.T) {
	query := fmt.Sprintf("insert into user_friendships (%s) values (?,?)", insertRemark)
	scan, err := sqlDriver.ExecSql(query, 1111111, "2020-01-11 19:42:22")
	if err != nil {
		t.Errorf("insert fail %v", err)
		return
	}

	m, err := scan.scanner()
	if err != nil {
		t.Errorf("scanner fail %v", err)
		return
	}

	fmt.Printf("last result %d", m)

}

func TestMysql_Transaction(t *testing.T) {
	err := sqlDriver.Transaction(func(tx interface{}) error {
		iq := fmt.Sprintf("insert into user_friendships (%s) values (?,?)", insertRemark)
		scan, err := tx.(*sqlx.Tx).ExecContext(context.Background(), iq, 1111111, "2020-01-11 19:42:22")
		if err != nil {
			return err
		}
		m, err := scan.RowsAffected()
		if err != nil {
			t.Errorf("scanner fail %v", err)
			return err
		}

		fmt.Printf("last result %d", m)
		query := fmt.Sprintf("select %s from user_friendships where to_uid in (?)", userRemark)
		rows, err := tx.(*sqlx.Tx).QueryxContext(context.Background(), query, 1111111)
		if err != nil {
			return err
		}
		var c = map[string]interface{}{}
		for rows.Next() {
			if err = rows.MapScan(c); err == nil {
				fmt.Printf("scan value %+v \n", c)
			}

		}
		return nil
	})

	if err != nil {
		t.Errorf("Transaction fail %+v", err)
		return
	}

}

func TestNewSqlx(t *testing.T) {
	freeConn := []uint32{1,3,2,4}
	closing := make([]uint32,0)
	for i := 0; i < len(freeConn); i++ {
		c := freeConn[i]
		closing = append(closing, c)
		last := len(freeConn) - 1
		freeConn[i] = freeConn[last]
		freeConn[last] = 0
		freeConn = freeConn[:last]
		i--
	}
}
