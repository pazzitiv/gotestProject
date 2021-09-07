package app

type Currency struct {
	Name  string `yaml:"name" json:"name"`
	Value int    `yaml:"value" json:"value"`
}

type Metrics struct {
	Currencies []Currency `yaml:"currencies" json:"currencies"`
}
