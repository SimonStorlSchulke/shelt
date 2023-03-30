package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func toAttributes() {

}

func GetArticle(c *gin.Context) {

	articleUrl := "articles/" + c.Param("id")
	templateUrl := "article-templates/" + c.Query("template-id")

	articleBody, err := StrapiGet(articleUrl + "?populate=deep")
	logErr(err)

	templateBody, err := StrapiGet(templateUrl)
	logErr(err)

	articleData := gjson.GetBytes(articleBody, "data.attributes").Value()
	animalData := gjson.GetBytes(articleBody, "data.attributes.Animals.data").Value().([]interface{})

	var an []interface{} = make([]interface{}, len(animalData))

	//de-nest Animal Data
	for i, animal := range animalData {
		an[i] = animal.(map[string]interface{})["attributes"]
	}

	adjustedArticleData := make(map[string]interface{})

	for k, r := range articleData.(map[string]interface{}) {
		if k != "Animals" {
			adjustedArticleData[k] = r
		}
	}

	adjustedArticleData["Animals"] = an

	templateData := gjson.GetBytes(templateBody, "data.attributes").Map()

	html, err := RenderTemplate(templateData["Html"].String(), adjustedArticleData)

	logErr(err)

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}
