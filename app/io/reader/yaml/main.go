package yaml

import (
	"gopkg.in/yaml.v2"
	"gotestProject/app"
	"io/ioutil"
	"reflect"
)

type YamlConfig struct {
	FilePath string
}

type Reader struct {
}

func (y Reader) Read(cfg interface{}) (app.Metrics, error) {
	var (
		config YamlConfig
		result app.Metrics
	)

	config.FilePath = reflect.ValueOf(cfg).FieldByName("FilePath").String()

	yamlFile, err := ioutil.ReadFile(config.FilePath)
	if err != nil {
		return app.Metrics{}, err
	}
	err = yaml.Unmarshal(yamlFile, &result)
	if err != nil {
		return app.Metrics{}, err
	}

	return result, nil
}
