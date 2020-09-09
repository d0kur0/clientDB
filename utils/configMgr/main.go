package configMgr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

var config = ServerConfig{
	Port: defaultPort,
}

func Get() ServerConfig {
	return config
}

func Parse() (err error) {
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		jsonBytes, err := json.MarshalIndent(config, "", "\t")
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Marshaling config error (%s)", configFilePath))
		}

		if err := ioutil.WriteFile(configFilePath, jsonBytes, 0644); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Write in config data file error (%s)", configFilePath))
		}
	} else {
		jsonData, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Read config file error (%s)", configFilePath))
		}

		if err := json.Unmarshal(jsonData, &config); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Unmarshal config file error (%s)", configFilePath))
		}
	}

	return
}
