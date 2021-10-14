package parser

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	sqlFile := filepath.Join("/home/lalifeier/project/go-gen-tools/model/sql/example/sql", "user.sql")
	fmt.Println(sqlFile)
	fmt.Println("============")
	tables, err := Parse(sqlFile)
	fmt.Println(tables)
	assert.Nil(t, err)
}
