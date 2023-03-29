package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
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

	animalCollectionData := gjson.GetBytes(animalCollectionBody, "data").Array()
	templateData := gjson.GetBytes(templateBody, "data.attributes").Map()

	var animalList []interface{}

	for _, animalStrapiData := range animalCollectionData {
		animalData := animalStrapiData.Map()["attributes"]
		animalList = append(animalList, animalData.Value())

		logErr(err)
	}
	html, err := RenderTemplate(templateData["Html"].String(), animalList)

	logErr(err)

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}
