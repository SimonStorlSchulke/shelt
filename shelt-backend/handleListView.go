package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAnimalListView(c *gin.Context) {

	templateUrl := "list-templates/" + c.Query("template-id")
	animalsUrl := "animals?populate=*"

	tagsString := c.Query("tags")

	if tagsString != "" {
		tags := strings.Split(tagsString, ",")

		for i, tag := range tags {
			animalsUrl += fmt.Sprintf("&filters[Tags][Name][[$in][%v]=%s", i, tag)
		}
	}

	animalCollectionBody, err := StrapiGet(animalsUrl)
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

	var animalList []map[string]interface{}

	for _, animalStrapiData := range animalCollectionData {
		animalData := animalStrapiData.(map[string]interface{})["attributes"]
		animalList = append(animalList, animalData.(map[string]interface{}))

		logErr(err)
	}
	html, err := RenderTemplate(templateData["Html"].(string), animalList)

	logErr(err)

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}
