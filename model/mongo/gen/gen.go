package gen

import (
	"path/filepath"

	"github.com/lalifeier/go-gen-tools/model/mongo/template"
	"github.com/lalifeier/go-gen-tools/util/file"
	"github.com/lalifeier/go-gen-tools/util/templatex"
)

const (
	category          = "mongo"
	modelTemplateFile = "model.tpl"
	errTemplateFile   = "err.tpl"
)

type Options struct {
	Types  []string
	Output string
}

func Parse(opts *Options) error {
	err := generateModel(opts)
	if err != nil {
		return err
	}
	return nil
}

func generateModel(opts *Options) error {
	var filename = ""
	for _, t := range opts.Types {
		output := filepath.Join(opts.Output, filename+".go")

		text, err := file.LoadTemplate(category, modelTemplateFile, template.Text)
		if err != nil {
			return err
		}

		err = templatex.New("model").Parse(text).Format(true).SaveTo(map[string]interface{}{
			"Type": t,
		}, output, false)
		if err != nil {
			return err
		}
	}
	return nil
}
