package app

type Currency struct {
	Name  string `yaml:"name" json:"name"`
	Value int    `yaml:"value" json:"value"`
}

type Currencies struct {
	Currencies []Currency `yaml:"currencies" json:"currencies"`
}
