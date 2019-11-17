package dao

import (
	"database/sql"
)

func Connect() (*sql.DB, error) {
	mysqlDB, err := connectMysql()
	if err != nil {
		return nil, err
	}
	return mysqlDB, nil
}
