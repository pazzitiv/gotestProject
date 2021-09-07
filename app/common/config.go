package common

import "gotestProject/app/io/reader/yaml"

type config struct {
	Server struct {
		Host string
		Port int
	}
	Readers struct {
		Yaml yaml.YamlConfig
	}
	Writers struct{}
}

func (c *config) Default() {
	c.Server.Host = "0.0.0.0"
	c.Server.Port = 80

	c.Readers.Yaml.FilePath = "data.yml"
}

var (
	App config
)
