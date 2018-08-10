package src

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

const assetsDir = "assets"
const stringsDir = "strings"
const spriteSheetsDir = "spritesheets"

var subDirs = []string{stringsDir, spriteSheetsDir}

const templatesDir = "src/templates" // TODO: hardcoded, use templates and run with makefile instead
const stringTemplateFile = "string.tpl"

type StringTplData struct {
	Description string
	Items       map[string]string
}

// mkdir creates directory under current working directory
// if doesn't exist
func mkdir(name string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	return os.Mkdir(path.Join(pwd, name), os.ModePerm)
}

func createFile(filePath string) (*os.File, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return os.Create(path.Join(pwd, filePath))
}

// CreateStringAsset generates string asset in strings directory with provided locale
// and (optionally) with provided items
func CreateStringAsset(name string, locale string, data interface{}) {
	rawTpl, err := ioutil.ReadFile(path.Join(templatesDir, stringTemplateFile))
	if err != nil {
		fmt.Printf("Fatal: failed to open template file %s\n", stringTemplateFile)
		fmt.Println(err)
		return
	}
	strTpl := string(rawTpl)
	tpl := template.New("String resource template")
	tpl, err = tpl.Parse(strTpl)
	if err != nil {
		fmt.Printf("Fatal: failed to parse template file %s\n", stringTemplateFile)
		fmt.Println(err)
		return
	}
	strFile, err := createFile(path.Join(assetsDir, stringsDir, locale, name))
	if err != nil {
		fmt.Printf("Fatal: failed to create strings file %s\n", name)
		fmt.Println(err)
		return
	}
	if data == nil {
		dataExample := StringTplData{
			Description: "Example string resource",
			Items: map[string]string{
				"foo": "bar",
				"go":  "lang",
			},
		}
		err = tpl.Execute(strFile, dataExample)
	} else {
		err = tpl.Execute(strFile, data)
	}
	if err != nil {
		fmt.Printf("Fatal: failed to generate strings file %s\n", name)
		fmt.Println(err)
		return
	}
}

// CreateAssets generates assets directory with example assets
func CreateAssets() {
	err := mkdir(assetsDir)
	if err != nil {
		fmt.Println("Fatal: failed to create assets dir")
		fmt.Println(err)
		return
	}
	for _, dir := range subDirs {
		err := mkdir(path.Join(assetsDir, dir))
		if err != nil {
			fmt.Printf("Fatal: failed to create %s dir\n", dir)
			fmt.Println(err)
			return
		}
	}
	for _, localeDir := range localesAll {
		err := mkdir(path.Join(assetsDir, stringsDir, localeDir))
		if err != nil {
			fmt.Printf("Fatal: failed to create string locales %s dir\n", localeDir)
			fmt.Println(err)
			return
		}
	}
	// Generate examples
	CreateStringAsset("example.yaml", LcEN, nil)
}
