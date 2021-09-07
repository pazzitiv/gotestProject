package reader

import "gotestProject/app"

type Reader interface {
	Read(interface{}) (app.Currencies, error)
}

type Readers struct {
	reader Reader
}

func (r *Readers) SetReader(reader Reader) {
	r.reader = reader
}

func (r *Readers) GetReader() Reader {
	return r.reader
}
