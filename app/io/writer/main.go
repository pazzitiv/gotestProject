package writer

import "gotestProject/app"

type Writer interface {
	Write(currencies app.Currencies) ([]byte, error)
}

type Writers struct {
	writer Writer
}

func (r *Writers) SetWriter(writer Writer) {
	r.writer = writer
}

func (r *Writers) GetWriter() Writer {
	return r.writer
}
