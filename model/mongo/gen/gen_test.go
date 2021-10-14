package gen

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/lalifeier/go-gen-tools/util/file"
	"github.com/stretchr/testify/assert"
)

var testTypes = `
	type User struct{}
	type Class struct{}
`

func TestGen(t *testing.T) {
	tempDir := file.MustTempDir()

	typesfile := filepath.Join(tempDir, "types.go")
	err := ioutil.WriteFile(typesfile, []byte(testTypes), 0o666)
	assert.Nil(t, err)

	t.Log(typesfile)

	err = Parse(&Options{
		Types:  []string{"User", "Class"},
		Output: tempDir,
	})

	assert.Nil(t, err)
}
