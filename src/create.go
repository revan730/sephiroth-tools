package src

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"text/template"
)

const assetsDir = "assets"
const stringsDir = "strings"
const spriteSheetsDir = "spritesheets"
const fontsDir = "fonts"

var subDirs = []string{stringsDir, spriteSheetsDir, fontsDir}

const templatesDir = "templates"
const stringTemplateFile = "string.tpl"

type StringAsset struct {
	Description string
	Items       map[string]string
}

func logError(msg string, err error) {
	fmt.Println("Fatal: ", msg)
	fmt.Println(err)
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

func writeStringsFile(wg *sync.WaitGroup, tpl *template.Template, name string, locale string, data interface{}) {
	defer wg.Done()

	strFile, err := createFile(path.Join(assetsDir, stringsDir, locale, name))
	if err != nil {
		logError("failed to create strings file "+name, err)
		return
	}
	if data == nil {
		dataExample := StringAsset{
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
		logError("failed to generate strings file "+name, err)
		return
	}
}

// CreateStringAsset generates string asset in strings directory for each supported locale
// and (optionally) with provided items
func CreateStringAsset(name string, data interface{}) {
	rawTpl, err := ioutil.ReadFile(path.Join(templatesDir, stringTemplateFile))
	if err != nil {
		logError("failed to open template file "+stringTemplateFile, err)
		return
	}
	strTpl := string(rawTpl)
	tpl := template.New("String resource template")
	tpl, err = tpl.Parse(strTpl)
	if err != nil {
		logError("failed to parse template file "+stringTemplateFile, err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(localesAll))
	for _, locale := range localesAll {
		go writeStringsFile(&wg, tpl, name, locale, data)
	}
	wg.Wait()
}

// CreateAssets generates assets directory with example assets
func CreateAssets() {
	err := mkdir(assetsDir)
	if err != nil {
		logError("failed to create assets dir", err)
		return
	}
	for _, dir := range subDirs {
		err := mkdir(path.Join(assetsDir, dir))
		if err != nil {
			logError("failed to create dir "+dir, err)
			return
		}
	}
	for _, localeDir := range localesAll {
		err := mkdir(path.Join(assetsDir, stringsDir, localeDir))
		if err != nil {
			logError("failed to create string locales dir "+localeDir, err)
			return
		}
	}
	// Generate examples
	CreateStringAsset("example.yaml", nil)
}
