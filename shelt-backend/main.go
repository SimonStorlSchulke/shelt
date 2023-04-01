package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func logErr(err error) {
	if err != nil {
		log.Println("ERROR:", err)
	}
}

type M map[string]interface{}

var strapiKey string = os.Args[1]

var defaultTemplates *template.Template

func init() {

	var err error
	defaultTemplates, _ = template.New("defaulttemplates").Funcs(templateFuncs).ParseGlob("./template-defaults/*.html")
	logErr(err)

	recollectArticleIds()
	go func() {
		for range time.Tick(time.Second * 30) {
			recollectArticleIds()
		}
	}()
}

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.GET("/animal/:id", GetAnimalFieldsById)
	r.GET("/animals", GetAnimalCollection)
	r.GET("/article/:id", GetArticle)
	r.GET("/animal-view", GetAnimalView)
	r.GET("/animal-list/:type", GetAnimalListView)
	r.GET("/uploads/:route", strapiUploads)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

var animalArticleIds map[string]int = make(map[string]int)

func recollectArticleIds() {
	fmt.Println("recollect article links")

	data, err := StrapiGet("animals?fields[0]=Name&fields[1]=CustumArticleLink&populate=Article&pagination[pageSize]=10000")
	if err != nil {
	}

	animals := gjson.GetBytes(data, "data").Array()

	for _, animal := range animals {
		animalName := gjson.Get(animal.Raw, "attributes.Name").String()
		articleId := gjson.Get(animal.Raw, "attributes.Article.data.id").Value()
		if articleId != nil {
			animalArticleIds[animalName] = int(math.Floor(articleId.(float64)))
		}
	}
}
