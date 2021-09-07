package yaml

import (
	"gopkg.in/yaml.v2"
	"gotestProject/app"
	"io/ioutil"
)

type YamlConfig struct {
	FilePath string
}

type Reader struct {
}

func (y Reader) Read(filepath string) (app.Currencies, error) {
	var (
		result app.Currencies
	)

	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return app.Currencies{}, err
	}
	err = yaml.Unmarshal(yamlFile, &result)
	if err != nil {
		return app.Currencies{}, err
	}

	return result, nil
}
