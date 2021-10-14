package templatex

import (
	"bytes"
	"go/format"
	"html/template"
	"io/ioutil"

	"github.com/lalifeier/go-gen-tools/util/file"
)

const regularPerm = 0o666

type DefaultTemplate struct {
	name   string
	text   string
	format bool
}

func New(name string) *DefaultTemplate {
	return &DefaultTemplate{
		name: name,
	}
}

func (t *DefaultTemplate) Parse(text string) *DefaultTemplate {
	t.text = text
	return t
}

func (t *DefaultTemplate) Format(format bool) *DefaultTemplate {
	t.format = format
	return t
}

func (t *DefaultTemplate) SaveTo(data interface{}, path string, forceUpdate bool) error {
	if file.Exists(path) && !forceUpdate {
		return nil
	}

	output, err := t.Execute(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, output.Bytes(), regularPerm)
}

func (t *DefaultTemplate) Execute(data interface{}) (*bytes.Buffer, error) {
	tpl, err := template.New(t.name).Parse(t.text)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err = tpl.Execute(buf, data); err != nil {
		return nil, err
	}

	if !t.format {
		return buf, nil
	}

	formatOutput, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}

	buf.Reset()
	buf.Write(formatOutput)
	return buf, nil
}
