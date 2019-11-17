package config

import (
	"errors"
)

func GetMysqlConfig() (map[interface{}]interface{}, error) {
	config, err := loadYAML(globalConfigFile)
	if err != nil {
		return nil, err
	}

	mysqlConfig, exist := config["mysql"]
	if exist {
		mysqlConfigMap, ok := mysqlConfig.(map[interface{}]interface{})
		if !ok {
			return nil, errors.New("too few values of 'mysql' at config file")
		}

		if !mysqlConfigMap["enabled"].(bool) {
			return nil, errors.New("mysql is disabled at config file")
		}

		return mysqlConfigMap, nil
	} else {
		return nil, errors.New("config file has not key 'mysql'")
	}
}
