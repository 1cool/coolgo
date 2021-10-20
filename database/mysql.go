package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"time"
)

const (
	SqlConnError = "SQL connect error."
	SqlPingError = "SQL ping error."
)

type Mysql struct {
	Host     string
	Port     int
	User     string
	Pwd      string
	Database string
}

func Instance(cfg Mysql) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.User, cfg.Pwd, cfg.Host, cfg.Port, cfg.Database)

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return db, errors.Wrap(err, SqlConnError)
	}

	err = db.Ping()

	if err != nil {
		return db, errors.Wrap(err, SqlPingError)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
