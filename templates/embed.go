package templates

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func BindFunctions(templ *template.Template) {
	funcs := template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
	}

	templ.Funcs(funcs)
}

func BindTemplates(templ *template.Template) error {
	return filepath.Walk("templates/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".tmpl" {
			templateName := strings.TrimSuffix(filepath.Base(path), ".tmpl")
			baseFolder := strings.TrimPrefix(filepath.Dir(path), "templates")

			if len(baseFolder) > 0 {
				templateName = strings.TrimPrefix(baseFolder, "/") + "/" + templateName
			}

			templateContent, _ := os.ReadFile(path)

			if _, err := templ.New(templateName).Parse(string(templateContent)); err != nil {
				panic(err)
			}
		}

		return nil
	})
}
