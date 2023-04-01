package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func GetAnimalView(c *gin.Context) {

	animalUrl := "animals/" + c.Query("animal-id")
	templateUrl := "templates/" + c.Query("template-id")

	animalBody, err := StrapiGet(animalUrl + "?populate=*")
	logErr(err)

	templateBody, err := StrapiGet(templateUrl)
	logErr(err)

	animalData := gjson.GetBytes(animalBody, "data.attributes").Value()
	templateData := gjson.GetBytes(templateBody, "data.attributes").Map()

	html, err := RenderTemplate(templateData["Html"].String(), animalData)

	logErr(err)

	c.Data(http.StatusOK, "text/html; charset=utf-8", html)
}
