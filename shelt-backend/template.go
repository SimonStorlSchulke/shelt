package main

import (
	"bytes"
	"html/template"
)

// RenderTemplate executes a template on data and returns the fesulting html
func RenderTemplate(templateString string, data map[string]interface{}) (string, error) {

	t := template.New("action")

	t, err := t.Parse(templateString)
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
