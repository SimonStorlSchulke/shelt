package main

import (
	"fmt"
	"html/template"
	"math"
	"time"
)

// provided to templates in RenderTemplate()

var templateFuncs template.FuncMap = template.FuncMap{
	"Age":         Age,
	"ArticleLink": ArticleLink,
}

// CAREFUL: when using cached htmls, this will obviously become outdated. better use js in this case.
func Age(animal interface{}) string {
	birthDate := animal.(map[string]interface{})["BirthDate"]
	birthDay, err := time.Parse("2006-01-02", fmt.Sprintf("%v", birthDate))
	if err != nil {
		return " - "
	}
	age := time.Since(birthDay)

	days := age.Hours() / 24
	years := days / 365

	if years < 3 && math.Mod(years, 1) > 0.5 {
		return fmt.Sprintf("%v½", math.Floor(years))
	}

	return fmt.Sprintf("%v", math.Floor(years))
}

func ArticleLink(animal interface{}) string {
	articleData := animal.(map[string]interface{})["Article"].(map[string]interface{})["data"]

	if articleData != nil {
		fmt.Println(animal)
		id := int(articleData.(map[string]interface{})["id"].(float64))

		return fmt.Sprintf("/article/%v", id)
	}
	return ""
}
