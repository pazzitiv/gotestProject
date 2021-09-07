package web

import (
	"gotestProject/app"
	"gotestProject/app/common"
	readers "gotestProject/app/io/reader"
	yamlreader "gotestProject/app/io/reader/yaml"
	writers "gotestProject/app/io/writer"
	jsonreader "gotestProject/app/io/writer/json"
	omwriter "gotestProject/app/io/writer/openmetrics"
	"log"
	"net/http"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		config         interface{}
		err            error
		source, format string
		data           app.Currencies
		result         []byte
		rd             readers.Readers
		wr             writers.Writers
	)

	err = r.ParseForm()
	if err != nil {
		log.Printf("Request error: %s", err.Error())
	}

	source = r.Form.Get("source")
	format = r.Form.Get("format")

	/**
	Выбор формата ввода
	*/
	switch source {
	case "yaml":
		fallthrough
	default:
		config = common.App.Readers.Yaml
		rd.SetReader(yamlreader.Reader{})

	}

	data, err = rd.GetReader().Read(config)
	if err != nil {
		log.Printf("Reader error: %s", err.Error())
	}

	/**
	Выбор формата вывода
	*/
	switch format {
	case "json":
		wr.SetWriter(jsonreader.JSONWriter{})

	case "openmetrics":
		fallthrough
	default:
		wr.SetWriter(omwriter.OMWriter{})

	}

	result, err = wr.GetWriter().Write(data)
	if err != nil {
		log.Fatalf("[FATAL] Write error: %s", err.Error())
	}

	_, err = w.Write(result)
	if err != nil {
		log.Fatalf("[FATAL] Write response error: %s", err.Error())
	}
}
