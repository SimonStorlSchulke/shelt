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

	articleBody, err := StrapiGet(articleUrl + "?populate=deep,3")
	if err != nil {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<p>505 Bad Gateway</p>"))
		return
	}

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

	templateId := c.Query("template-id")

	var html []byte

	if templateId != "" {
		templateUrl := "article-templates/" + templateId
		templateBody, err := StrapiGet(templateUrl)
		logErr(err)
		templateData := gjson.GetBytes(templateBody, "data.attributes").Map()
		html, err = RenderTemplate(templateData["Html"].String(), adjustedArticleData)
	} else {
		html, err = RenderTemplateFile("article-default.html", adjustedArticleData)
	}

	logErr(err)

	addArticleLinks(&html)

	c.Data(http.StatusOK, "text/html; charset=utf-8", html)
}
