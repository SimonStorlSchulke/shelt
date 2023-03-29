package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {

	animalUrl := "animals/" + c.Query("animal-id")
	templateUrl := "templates/" + c.Query("template-id")

	animalBody, err := StrapiGet(animalUrl + "?populate=*")
	logErr(err)

	templateBody, err := StrapiGet(templateUrl)
	logErr(err)

	var animalStrapiData gin.H
	var templateStrapiData gin.H

	json.Unmarshal(animalBody, &animalStrapiData)
	json.Unmarshal(templateBody, &templateStrapiData)

	animalData := animalStrapiData["data"].(gin.H)["attributes"]
	templateData := templateStrapiData["data"].(gin.H)["attributes"].(gin.H)

	html, err := RenderTemplate(templateData["Html"].(string), animalData.(gin.H))

	logErr(err)

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}
