package main

import (
	"bytes"
	"html/template"
)

// RenderTemplate executes a template on data and returns the fesulting html
func RenderTemplate(templateString string, data any) (string, error) {

	t, err := template.New("action").Funcs(template.FuncMap{
		"TestTemplateFunc": TestTemplateFunc,
		"AgeFromBirthDate": AgeFromBirthDate,
	}).Parse(templateString)

	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}

	result := tpl.String()

	return result, nil
}
