package dao

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"project-support-system/config"
)

func connectMysql() (*sql.DB, error) {
	mysqlConfig, err := config.GetMysqlConfig()
	if err != nil {
		return nil, err
	}

	var mysqlDB *sql.DB
	mysqlDB, err = sql.Open("mysql", mysqlConfig["url"])
	if err != nil {
		return nil, err
	}

	mysqlDB.SetMaxIdleConns(mysqlConfig["maxIdleConnection"].(int))
	mysqlDB.SetMaxOpenConns(mysqlConfig["maxOpenConnection"].(int))

	if err = mysqlDB.Ping(); err != nil {
		return nil, err
	}

	return mysqlDB, nil
}
