package json

import (
	"encoding/json"
	"gotestProject/app"
)

type Label struct {
	Name  string
	Value string
}

type JSONWriter struct {
}

func (w JSONWriter) Write(data app.Currencies) ([]byte, error) {
	return json.Marshal(data)
}
