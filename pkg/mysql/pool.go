package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // msyql driver
	"github.com/jmoiron/sqlx"
)

// Pool 新建Mysql pool
func Pool(addr string, port int, user, passwd, dbName, charSet string,
	maxOpen, maxIdle int) (*sqlx.DB, error) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		user, passwd, addr, port, dbName, charSet)
	pool, err := sqlx.Connect("mysql", source)
	if err != nil {
		return nil, err
	}
	pool.SetMaxOpenConns(maxOpen)
	pool.SetMaxIdleConns(maxIdle)
	return pool, nil
}
