package util

import (
	"fmt"
)

var goKeyword = map[string]string{
	"var":         "variable",
	"const":       "constant",
	"package":     "pkg",
	"func":        "function",
	"return":      "rtn",
	"defer":       "dfr",
	"go":          "goo",
	"select":      "slt",
	"struct":      "structure",
	"interface":   "itf",
	"chan":        "channel",
	"type":        "tp",
	"map":         "mp",
	"range":       "rg",
	"break":       "brk",
	"case":        "caz",
	"continue":    "ctn",
	"for":         "fr",
	"fallthrough": "fth",
	"else":        "es",
	"if":          "ef",
	"switch":      "swt",
	"goto":        "gt",
	"default":     "dft",
}

func EscapeGolangKeyword(s string) string {
	if !isGolangKeyword(s) {
		return s
	}

	r := goKeyword[s]
	fmt.Printf("[EscapeGolangKeyword]: go keyword is forbidden %q, converted into %q", s, r)
	return r
}

func isGolangKeyword(s string) bool {
	_, ok := goKeyword[s]
	return ok
}
