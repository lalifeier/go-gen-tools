package file

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	NL       = "\n"
	goctlDir = ".goctl"
)

var goctlHome string

func MustTempDir() string {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatalln(err)
	}

	return dir
}

func Exists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func GetGoctlHome() (string, error) {
	if len(goctlHome) != 0 {
		return goctlHome, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, goctlDir), nil
}

func GetTemplateDir(category string) (string, error) {
	goctlHome, err := GetGoctlHome()
	if err != nil {
		return "", err
	}

	return filepath.Join(goctlHome, category), nil
}

func LoadTemplate(category, file, text string) (string, error) {
	dir, err := GetTemplateDir(category)
	if err != nil {
		return "", err
	}

	file = filepath.Join(dir, file)
	if !Exists(file) {
		return text, nil
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
