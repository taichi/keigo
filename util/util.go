package util

import (
	"bytes"
	"github.com/taichi/keigo/log"
	"sync"
	"text/template"
)

func ToOnceFn(f func()) func() {
	var once sync.Once
	return func() {
		once.Do(f)
	}
}

func MaybeFault(err error) {
	if err != nil {
		log.Panic(err)
	}
}

type MustTemplate struct {
	delegate *template.Template
}

func Must(name, text string) *MustTemplate {
	tmpl := template.Must(template.New(name).Parse(text))
	return &MustTemplate{delegate: tmpl}
}

func (mt *MustTemplate) Do(arg interface{}) string {
	var buf bytes.Buffer
	err := mt.delegate.Execute(&buf, arg)
	MaybeFault(err)
	return buf.String()
}
