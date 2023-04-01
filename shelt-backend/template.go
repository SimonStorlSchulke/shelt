package main

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
)

var templateFuncs template.FuncMap = template.FuncMap{
	"TestTemplateFunc": TestTemplateFunc,
	"AgeFromBirthDate": AgeFromBirthDate,
}

func addArticleLinks(html *bytes.Buffer) *bytes.Buffer {
	rx := regexp.MustCompile(`\(\(.*?\)\)`)

	matches := rx.FindAll(html.Bytes(), -1)

	var b []byte = html.Bytes()

	for _, match := range matches {

		animalName := string(match)
		animalName = animalName[2 : len(animalName)-2] // "((animalname))"" --> "animalname"

		if id, ok := animalArticleIds[animalName]; ok {
			link := []byte(fmt.Sprintf(`<a href="/article/%v">%s</a>`, id, animalName))
			b = bytes.ReplaceAll(b, match, link)
		} else {
			b = bytes.ReplaceAll(b, match, []byte(animalName))
		}
	}
	return bytes.NewBuffer(b)
}

// RenderTemplate executes a template on data and returns the fesulting html
func RenderTemplate(templateString string, data any) (bytes.Buffer, error) {
	t, err := template.New("action").Funcs(templateFuncs).Parse(templateString)

	var tpl bytes.Buffer
	if err != nil {
		return tpl, err
	}
	//data = addArticleLinks(data)
	if err := t.Execute(&tpl, data); err != nil {
		return tpl, err
	}

	//tpl = addDogLinks(&tpl)

	newH := addArticleLinks(&tpl)

	return *newH, nil
}

// RenderTemplateFile executes a template file on data and returns the fesulting html
func RenderTemplateFile(path string, data any) (bytes.Buffer, error) {

	t, err := template.ParseFiles(path)

	t.Funcs(template.FuncMap(templateFuncs))

	var tpl bytes.Buffer

	if err != nil {
		return tpl, err
	}

	if err := t.Execute(&tpl, data); err != nil {
		return tpl, err
	}

	newH := addArticleLinks(&tpl)

	return *newH, nil
}
