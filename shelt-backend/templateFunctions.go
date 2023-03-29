package main

import (
	"fmt"
	"math"
	"time"
)

// provided to templates in RenderTemplate()

func TestTemplateFunc(p1 string) string {
	return "HELLO TEMPLATE FUNCTION" + p1
}

// CAREFUL: when using cached htmls, this will obviously become outdated. better use js in this case.
func AgeFromBirthDate(birthDayString interface{}) string {
	birthDay, err := time.Parse("2006-01-02", fmt.Sprintf("%v", birthDayString))
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
