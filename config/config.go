package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	globalConfigFile string
)

func Init(configFile string) error {
	globalConfigFile = configFile

	_, err := GetMysqlConfig()
	if err != nil {
		return err
	}

	return nil
}

func loadYAML(configFile string) (map[string]interface{}, error) {
	configFileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	conf := make(map[string]interface{})

	err = yaml.Unmarshal([]byte(configFileContent), &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
