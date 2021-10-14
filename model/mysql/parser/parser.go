package parser

import (
	"fmt"

	"github.com/lalifeier/go-gen-tools/util"
	"github.com/zeromicro/ddl-parser/parser"
)

type (
	// Table describes a mysql table
	Table struct {
		Name        string
		PrimaryKey  Primary
		UniqueIndex map[string][]*Field
		Fields      []*Field
	}

	// Primary describes a primary key
	Primary struct {
		Field
		AutoIncrement bool
	}

	// Field describes a table field
	Field struct {
		Name    string
		Type    string
		Comment string
	}
)

func Parse(filename string) ([]*Table, error) {
	p := parser.NewParser()
	ts, err := p.From(filename)
	if err != nil {
		return nil, err
	}
	tables := GetSafeTables(ts)
	for _, e := range tables {
		columns := e.Columns

		fmt.Println("============")
		fmt.Println(columns)
	}
	return nil, nil
}

func GetSafeTables(tables []*parser.Table) []*parser.Table {
	var list []*parser.Table
	for _, t := range tables {
		table := GetSafeTable(t)
		list = append(list, table)
	}

	return list
}

func GetSafeTable(table *parser.Table) *parser.Table {
	table.Name = util.EscapeGolangKeyword(table.Name)
	for _, c := range table.Columns {
		c.Name = util.EscapeGolangKeyword(c.Name)
	}

	for _, e := range table.Constraints {
		var uniqueKeys, primaryKeys []string
		for _, u := range e.ColumnUniqueKey {
			uniqueKeys = append(uniqueKeys, util.EscapeGolangKeyword(u))
		}
		for _, p := range e.ColumnPrimaryKey {
			primaryKeys = append(primaryKeys, util.EscapeGolangKeyword(p))
		}
		e.ColumnUniqueKey = uniqueKeys
		e.ColumnPrimaryKey = primaryKeys
	}
	return table
}
