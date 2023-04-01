package main

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
)

func addArticleLinks(json *[]byte) {
	rx := regexp.MustCompile(`\(\(.*?\)\)`)

	matches := rx.FindAll(*json, -1)

	for _, match := range matches {

		animalName := string(match)
		animalName = animalName[2 : len(animalName)-2] // "((animalname))"" --> "animalname"

		if id, ok := animalArticleIds[animalName]; ok {
			link := []byte(fmt.Sprintf(`<a href="/article/%v">%s</a>`, id, animalName))
			*json = bytes.ReplaceAll(*json, match, link)
		} else {
			*json = bytes.ReplaceAll(*json, match, []byte(animalName))
		}
	}
}

// RenderTemplate executes a template on data and returns the fesulting html
func RenderTemplate(templateString string, data any) ([]byte, error) {
	t, err := template.New("action").Funcs(templateFuncs).Parse(templateString)

	var tpl bytes.Buffer
	if err != nil {
		return tpl.Bytes(), err
	}

	if err := t.Execute(&tpl, data); err != nil {
		return tpl.Bytes(), err
	}

	return tpl.Bytes(), nil
}

// RenderTemplateFile executes a template file on data and returns the fesulting html
func RenderTemplateFileOld(path string, data any) ([]byte, error) {

	var tpl bytes.Buffer

	t, err := template.ParseFiles(path)
	if err != nil {
		return tpl.Bytes(), err
	}

	t.Funcs(template.FuncMap(templateFuncs))

	if err != nil {
		return tpl.Bytes(), err
	}

	if err := t.Execute(&tpl, data); err != nil {
		logErr(err)
		return tpl.Bytes(), err
	}

	return tpl.Bytes(), nil
}

// RenderTemplateFile executes a template file on data and returns the fesulting html
func RenderTemplateFile(filePath string, data any) ([]byte, error) {
	//not necessary in prod
	defaultTemplates, _ = template.New("defaulttemplates").Funcs(templateFuncs).ParseGlob("./template-defaults/*.html")

	var tpl bytes.Buffer

	if err := defaultTemplates.ExecuteTemplate(&tpl, filePath, data); err != nil {
		fmt.Println("GO SO MUCH", err)
		logErr(err)
		return tpl.Bytes(), err
	}

	return tpl.Bytes(), nil
}
