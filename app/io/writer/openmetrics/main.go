package openmetrics

import (
	"fmt"
	"gotestProject/app"
	"reflect"
	"strings"
)

type Label struct {
	Name  string
	Value string
}

type OMWriter struct {
}

func (w OMWriter) Write(data app.Metrics) ([]byte, error) {
	var (
		metricsResponse []string
	)

	om := reflect.ValueOf(data)
	omt := reflect.TypeOf(data)

	countMetrics := om.NumField()

	for i := 0; i < countMetrics; i++ {
		metric := om.Field(i)
		metricType := omt.Field(i)
		metricsResponse = append(metricsResponse, fmt.Sprintf("%s{%s}", metricType.Name, w.parseValue(metric)))
	}

	return []byte(strings.Join(metricsResponse, "\n")), nil
}

func (w OMWriter) parseValue(value reflect.Value) string {
	var (
		labelsString []string
	)

	labels := value.Interface().([]app.Currency)

	for _, lbl := range labels {
		labelsString = append(labelsString, fmt.Sprintf("Currency_%s=%v", lbl.Name, lbl.Value))
	}

	return strings.Join(labelsString, ",")
}
