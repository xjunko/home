package templates

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func TemplateParseFolder(tmpl *template.Template, path string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".tmpl" {
			_, err = tmpl.ParseFiles(path)
			if err != nil {
				fmt.Println(fmt.Errorf("failed to parse template: %v", err))
			}
		}

		return err
	})

}

func ParseTemplates(name string) *template.Template {
	funcs := template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
	}

	templ := template.New(name).Funcs(funcs)

	TemplateParseFolder(templ, "web/templates/html")
	TemplateParseFolder(templ, "web/templates/rss")

	return templ
}
