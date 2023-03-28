package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAnimalView(c *gin.Context) {
	//templateId := c.Param("template-id")

	animalUrl := "animals/" + c.Query("animal-id")
	templateUrl := "templates/" + c.Query("template-id")

	animalBody, err := StrapiGet(animalUrl + "?populate=*")
	logErr(err)

	templateBody, err := StrapiGet(templateUrl)
	logErr(err)

	var animalStrapiData map[string]interface{}
	var templateStrapiData map[string]interface{}

	json.Unmarshal(animalBody, &animalStrapiData)
	json.Unmarshal(templateBody, &templateStrapiData)

	animalData := animalStrapiData["data"].(map[string]interface{})["attributes"]
	templateData := templateStrapiData["data"].(map[string]interface{})["attributes"].(map[string]interface{})

	html, err := RenderTemplate(templateData["Html"].(string), animalData.(map[string]interface{}))

	logErr(err)

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

func GetAnimalCollectionView(c *gin.Context) {
	//templateId := c.Param("template-id")

	templateUrl := "templates/" + c.Query("template-id")

	//tags := c.Query("tags")

	animalCollectionBody, err := StrapiGet("animals?populate=*")
	logErr(err)

	templateBody, err := StrapiGet(templateUrl)
	logErr(err)

	var animalCollectionStrapiData map[string]interface{}
	var templateStrapiData map[string]interface{}

	err = json.Unmarshal(animalCollectionBody, &animalCollectionStrapiData)
	logErr(err)
	json.Unmarshal(templateBody, &templateStrapiData)

	animalCollectionData := animalCollectionStrapiData["data"].([]interface{})
	templateData := templateStrapiData["data"].(map[string]interface{})["attributes"].(map[string]interface{})

	fullHtml := ""
	for _, animalStrapiData := range animalCollectionData {
		animalData := animalStrapiData.(map[string]interface{})["attributes"]
		html, err := RenderTemplate(templateData["Html"].(string), animalData.(map[string]interface{}))
		logErr(err)
		fullHtml += html
	}

	logErr(err)

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fullHtml))
}
